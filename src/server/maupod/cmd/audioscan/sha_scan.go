package main

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"

	"github.com/mauleyzaola/maupod/src/server/pkg/broker"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func (m *MsgHandler) handlerSHAScan(msg *nats.Msg) {
	var input pb.MediaInfoInput
	var err error
	if err = proto.Unmarshal(msg.Data, &input); err != nil {
		m.base.Logger().Error(err)
		return
	}

	var timeout = time.Second * time.Duration(m.config.Delay)
	var filename string
	if input.Media != nil && input.Media.Location != "" {
		filename = input.Media.Location
	} else {
		filename = input.FileName
	}

	// read info from media in hd
	media, err := broker.RequestScanAudioFile(m.base.NATS(), m.base.Logger(), filename, timeout)
	if err != nil {
		m.base.Logger().Error(err)

		// if file is not readable, we asume we need to remove from db since it is invalid or no longer exists
		if err = broker.PublishMediaInfoDelete(m.base.NATS(), &input); err != nil {
			m.base.Logger().Error(err)
		}
		return
	}

	// calculate sha in file
	fileData, err := ioutil.ReadFile(media.Location)
	if err != nil {
		m.base.Logger().Error(err)
		return
	}

	shaStr := helpers.HashFromSHA(fileData)
	if shaStr == media.Sha {
		return
	}

	// update db
	media.ModifiedDate = helpers.TimeToTs2(time.Now())
	media.Sha = shaStr
	var cols = orm.MediumColumns
	var fields = []string{cols.Sha, cols.ModifiedDate}
	store := &data.MediaStore{}
	ctx := context.Background()
	conn := m.db
	if err = store.Update(ctx, conn, media, fields...); err != nil {
		m.base.Logger().Error(err)
		return
	}
}
