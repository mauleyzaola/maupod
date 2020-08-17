package main

import (
	"context"
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerMediaInfoDBSelect(msg *nats.Msg) {
	var err error
	var input pb.MediaInfoInput

	output := &pb.MediaInfoOutput{
		Response: &pb.Response{
			Ok:    false,
			Error: "",
		},
	}

	defer func() {
		if msg.Reply == "" {
			return
		}

		var localErr error
		var data []byte

		if data, localErr = helpers.ProtoMarshal(output); localErr != nil {
			log.Println(localErr)
			return
		}
		if localErr = msg.Respond(data); localErr != nil {
			log.Println(localErr)
			return
		}
	}()

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}

	if input.Media.Id == "" {
		output.Response.Ok = false
		output.Response.Error = "missing media.Id"
		return
	}

	ctx := context.Background()
	conn := m.db
	store := &dbdata.MediaStore{}
	mediaInfo, err := store.Select(ctx, conn, input.Media.Id)
	if err != nil {
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}
	output.Response.Ok = true
	output.Media = mediaInfo
	return
}
