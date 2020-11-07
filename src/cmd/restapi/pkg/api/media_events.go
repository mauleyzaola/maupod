package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mauleyzaola/maupod/src/protos"
	"io"
	"net/http"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
)

// MediaEventsGet will convert the current rows in table media_event to a special text file with this format, one JSON object for each line
// { JSON content }
// { JSON content }
func (a *ApiServer) MediaEventsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		conn := a.db
		rows, err := orm.MediaEvents().All(ctx, conn)
		if err != nil {
			helpers.WriteJson(w, err, http.StatusInternalServerError, nil)
			return
		}
		var filename = fmt.Sprintf("events-%s.txt", time.Now().Format(eventFileNameDateFormat))
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
		for _, row := range rows {
			event := MediaEvent(*row)
			if err = event.Write(w); err != nil {
				helpers.WriteJson(w, err, http.StatusInternalServerError, nil)
				return
			}
		}
	}
}

func (a *ApiServer) MediaEventsPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			helpers.WriteJson(w, err, http.StatusBadRequest, nil)
			return
		}
		defer func() {
			_ = file.Close()
		}()
		nc := a.nc
		var lineCount int
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var e *orm.MediaEvent
			line := scanner.Bytes()
			e, err = ParseMediaEvent(line)
			if err != nil {
				helpers.WriteJson(w, err, http.StatusBadRequest, nil)
				return
			}
			input := &protos.MediaEventInput{
				Id:    e.ID,
				Sha:   e.Sha,
				Ts:    helpers.TimeToTs(&e.TS),
				Event: protos.Message(e.Event),
			}
			if err = broker.PublishMediaEventUpsert(nc, input); err != nil {
				helpers.WriteJson(w, err, http.StatusInternalServerError, nil)
				return
			}
			lineCount++
		}

		helpers.WriteJson(w, nil, http.StatusAccepted, struct {
			Ok        bool `json:"ok"`
			LineCount int  `json:"line_count"`
		}{
			Ok:        true,
			LineCount: lineCount,
		})
	}
}

const eventFileNameDateFormat = "20060102-150405.123"

type MediaEvent orm.MediaEvent

func (e *MediaEvent) Write(w io.Writer) error {
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}
	if _, err = w.Write(data); err != nil {
		return err
	}
	if _, err = w.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}

func ParseMediaEvent(data []byte) (*orm.MediaEvent, error) {
	var event orm.MediaEvent
	if err := json.Unmarshal(data, &event); err != nil {
		return nil, err
	}
	return &event, nil
}

func ReadLineMediaEvent() {}
