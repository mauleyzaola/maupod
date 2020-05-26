package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/mauleyzaola/maupod/src/server/pkg/media"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerMediaInfo(msg *nats.Msg) {

	var err error
	var input pb.MediaInfoInput
	var output pb.MediaInfoOutput

	defer func() {
		var localErr error
		var data []byte

		if data, localErr = proto.Marshal(&output); localErr != nil {
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

	result, err := media.InfoFromFile(input.FileName)
	if err != nil {
		m.base.Logger().Error(err)
		output.Response.Ok = false
		output.Response.Error = err.Error()
		return
	}
	output.Media = result.ToProto()
	output.Response.Ok = true
	return
}
