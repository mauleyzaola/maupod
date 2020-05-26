package main

import (
	"context"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerAudioScan(msg *nats.Msg) {
	start := time.Now()
	var input pb.ScanDirectoryAudioFilesInput
	err := proto.Unmarshal(msg.Data, &input)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}

	defer func() {
		m.base.Logger().Info("[INFO] elapsed time: " + time.Since(start).String())
	}()

	m.base.Logger().Info("received artwork extract message: " + input.String())

	ctx := context.Background()
	conn := m.db

	if err = ScanDirectoryAudioFiles(
		ctx,
		conn,
		m.base.NATS(),
		m.base.Logger(),
		helpers.TsToTime2(input.ScanDate),
		&data.MediaStore{},
		input.Root,
		m.config,
	); err != nil {
		m.base.Logger().Error(err)
		return
	}
}
