package main

import (
	"os"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/information"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func (m *MsgHandler) handlerMediaInfo(msg *nats.Msg) {

	var err error
	var input pb.MediaInfoInput

	output := &pb.MediaInfoOutput{
		Response: &pb.Response{
			Ok:    false,
			Error: "",
		},
	}

	defer func() {
		var localErr error
		var data []byte

		if data, localErr = proto.Marshal(output); localErr != nil {
			m.base.Logger().Error(localErr)
			return
		}
		if localErr = msg.Respond(data); localErr != nil {
			m.base.Logger().Error(localErr)
			return
		}
	}()

	if err = proto.Unmarshal(msg.Data, &input); err != nil {
		m.base.Logger().Error(err)
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}
	m.base.Logger().Info("received media info message: " + input.String())

	result, err := information.MediaFromFile(input.FileName)
	if err != nil {
		m.base.Logger().Error(err)
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}

	info, err := os.Stat(input.FileName)
	if err != nil {
		m.base.Logger().Error(err)
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}
	output.LastModifiedDate = helpers.TimeToTs2(info.ModTime())
	output.Media = result
	output.Media.FolderName = filepath.Dir(input.FileName)
	output.Media.FileName = filepath.Base(input.FileName)
	output.Media.Location = filepath.Join(output.Media.FolderName, output.Media.FileName)
	output.Response.Ok = true

	return
}
