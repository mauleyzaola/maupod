package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/protos"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/information"
	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerMediaInfo(msg *nats.Msg) {
	var err error
	var input protos.MediaInfoInput

	output := &protos.MediaInfoOutput{
		Response: &protos.Response{
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
		log.Println(err)
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}
	log.Println("received media info message: " + input.String())

	var fullPath = input.FileName

	if _, err = os.Stat(fullPath); err != nil {
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}

	raw, err := information.MediaInfoFromFile(fullPath)
	if err != nil {
		log.Println(err)
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}
	var rawStr = raw.String()
	output.Raw = rawStr

	result, err := information.MediaFromRaw(rawStr)
	if err != nil {
		log.Println(err)
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}

	info, err := os.Stat(fullPath)
	if err != nil {
		log.Println(err)
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}
	var location = paths.LocationPath(fullPath)
	output.LastModifiedDate = helpers.TimeToTs2(info.ModTime())
	output.Media = result
	output.Media.FolderName = filepath.Dir(location)
	output.Media.FileName = filepath.Base(input.FileName)
	output.Media.Location = filepath.Join(output.Media.FolderName, output.Media.FileName)
	output.Response.Ok = true

	return
}
