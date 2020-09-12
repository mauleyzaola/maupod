package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
)

func (a *ApiServer) DistinctListGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var filter dbdata.MediaFilter
	if err = p.DecodeQuery(&filter); err != nil {
		status = http.StatusBadRequest
		return
	}
	filter.Distinct = p.Param("field")
	if result, err = a.mediaStore.DistinctList(p.ctx, p.conn, filter); err != nil {
		status = http.StatusBadRequest
		return
	}
	return
}

func (a *ApiServer) MediaListGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var filter dbdata.MediaFilter
	if err = p.DecodeQuery(&filter); err != nil {
		status = http.StatusBadRequest
		return
	}

	if result, err = a.mediaStore.List(p.ctx, p.conn, filter, nil); err != nil {
		status = http.StatusBadRequest
		return
	}
	return
}

func (a *ApiServer) AlbumViewListGet(p TransactionExecutorParams) (status int, result interface{}, err error) {
	var filter dbdata.MediaFilter
	if err = p.DecodeQuery(&filter); err != nil {
		status = http.StatusBadRequest
		return
	}

	if result, err = a.mediaStore.AlbumListView(p.ctx, p.conn, filter); err != nil {
		status = http.StatusBadRequest
		return
	}
	return
}

func (a *ApiServer) MediaSpectrumGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		conn := a.db
		id := mux.Vars(r)["id"]
		media, err := a.mediaStore.Select(ctx, conn, id)
		if err != nil {
			helpers.WriteJson(w, err, http.StatusNotFound, nil)
			return
		}
		height, err := strconv.Atoi(r.URL.Query().Get("height"))
		if err != nil {
			helpers.WriteJson(w, err, http.StatusBadRequest, nil)
			return
		}
		width, err := strconv.Atoi(r.URL.Query().Get("width"))
		if err != nil {
			helpers.WriteJson(w, err, http.StatusBadRequest, nil)
			return
		}

		var input = pb.SpectrumGenerateInput{
			Media:  media,
			Width:  int64(width),
			Height: int64(height),
		}
		var output pb.SpectrumGenerateOutput
		if err = broker.DoRequest(a.nc, pb.Message_MESSAGE_MEDIA_SPECTRUM_GENERATE, &input, &output, rules.Timeout(a.config)+(time.Second*5)); err != nil {
			log.Println(err)
			helpers.WriteJson(w, err, http.StatusInternalServerError, nil)
			return
		}
		if output.Error != "" {
			helpers.WriteJson(w, errors.New(output.Error), http.StatusInternalServerError, nil)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("content-type", "image/png")
		if _, err = w.Write(output.Data); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
}
