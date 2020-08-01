package main

import (
	"os"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/pkg/paths"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/information"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
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

		if data, localErr = helpers.ProtoMarshal(output); localErr != nil {
			m.base.Logger().Error(localErr)
			return
		}
		if localErr = msg.Respond(data); localErr != nil {
			m.base.Logger().Error(localErr)
			return
		}
	}()

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		m.base.Logger().Error(err)
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}
	m.base.Logger().Info("received media info message: " + input.String())

	var fullPath = paths.FullPath(input.FileName)
	var location = paths.LocationPath(fullPath)
	raw, err := information.MediaInfoFromFile(fullPath)
	if err != nil {
		m.base.Logger().Error(err)
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}
	var rawStr = raw.String()
	output.Raw = rawStr

	result, err := information.MediaFromRaw(rawStr)
	if err != nil {
		m.base.Logger().Error(err)
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}

	info, err := os.Stat(fullPath)
	if err != nil {
		m.base.Logger().Error(err)
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}
	output.LastModifiedDate = helpers.TimeToTs2(info.ModTime())
	output.Media = result
	output.Media.FolderName = filepath.Dir(location)
	output.Media.FileName = filepath.Base(input.FileName)
	output.Media.Location = filepath.Join(output.Media.FolderName, output.Media.FileName)
	output.Response.Ok = true

	return
}
