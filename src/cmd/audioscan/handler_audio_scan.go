package main

import (
	"context"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerAudioScan(msg *nats.Msg) {
	start := time.Now()
	var input pb.ScanDirectoryAudioFilesInput
	err := helpers.ProtoUnmarshal(msg.Data, &input)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}

	defer func() {
		m.base.Logger().Info("[INFO] elapsed time: " + time.Since(start).String())
	}()

	m.base.Logger().Info("received audio scan message: " + input.String())

	ctx := context.Background()
	conn := m.db

	if err = ScanDirectoryAudioFiles(
		ctx,
		conn,
		m.base.NATS(),
		m.base.Logger(),
		helpers.TsToTime2(input.ScanDate),
		&dbdata.MediaStore{},
		input.Root,
		m.config,
	); err != nil {
		m.base.Logger().Error(err)
		return
	}
}
