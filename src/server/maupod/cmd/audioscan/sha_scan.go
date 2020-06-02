package main

import (
	"context"
	"io/ioutil"

	"github.com/mauleyzaola/maupod/src/server/pkg/broker"
	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata"
	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata/orm"
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

	var filename string
	if input.Media != nil && input.Media.Location != "" {
		filename = input.Media.Location
	} else {
		filename = input.FileName
	}

	m.base.Logger().Info("xxx: " + filename)

	// read content from file system
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		m.base.Logger().Error(err)
		// if file is not readable, we asume we need to remove from db since it is invalid or no longer exists
		if err = broker.PublishMediaInfoDelete(m.base.NATS(), &input); err != nil {
			m.base.Logger().Error(err)
		}
		return
	}

	shaStr := helpers.HashFromSHA(fileData)
	if shaStr == input.Media.Sha {
		return
	}

	// update db
	input.Media.Sha = shaStr
	var cols = orm.MediumColumns
	store := &dbdata.MediaStore{}
	ctx := context.Background()
	conn := m.db
	if err = store.Update(ctx, conn, input.Media, cols.Sha); err != nil {
		m.base.Logger().Error(err)
		return
	}
}
