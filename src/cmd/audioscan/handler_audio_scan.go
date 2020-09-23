package main

import (
	"context"
	"log"
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
		log.Println(err)
		return
	}

	defer func() {
		log.Println("[INFO] elapsed time: " + time.Since(start).String())
	}()

	log.Println("received audio scan message: " + input.String())

	ctx := context.Background()
	conn := m.db

	if err = ScanDirectoryAudioFiles(
		ctx,
		conn,
		m.base.NATS(),
		helpers.TsToTime2(input.ScanDate),
		dbdata.NewMediaStore(),
		input.Root,
		m.config,
		input.Force,
	); err != nil {
		log.Println(err)
		return
	}
}
