package main

import (
	"context"
	"log"

	"github.com/mauleyzaola/maupod/src/protos"

	"github.com/mauleyzaola/maupod/src/pkg/types"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/nats-io/nats.go"
)

// handlerQueueList returns a list of queues along with the medias related
func (m *MsgHandler) handlerQueueList(msg *nats.Msg) {
	var output protos.QueueOutput

	defer func() {
		if msg.Reply == "" {
			return
		}
		data, err := helpers.ProtoMarshal(&output)
		if err != nil {
			log.Println(err)
			return
		}
		if err = msg.Respond(data); err != nil {
			log.Println(err)
		}
	}()
	ctx := context.Background()
	conn := m.db
	medias, err := queueList(ctx, conn)
	if err != nil {
		output.Error = err.Error()
		return
	}
	output.Rows = mediasToQueues(medias)
}

// handlerQueueAdd adds one media to the queue
func (m *MsgHandler) handlerQueueAdd(msg *nats.Msg) {
	var input protos.QueueInput
	var err error
	var output protos.QueueOutput

	defer func() {
		onQueueNotifyChanged(m.base.NATS())
		if msg.Reply == "" {
			return
		}
		data, err := helpers.ProtoMarshal(&output)
		if err != nil {
			log.Println(err)
			return
		}
		if err = msg.Respond(data); err != nil {
			log.Println(err)
		}
	}()

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		output.Error = err.Error()
		return
	}
	ctx := context.Background()
	conn := m.db
	list, err := queueList(ctx, conn)
	if err != nil {
		output.Error = err.Error()
		return
	}

	if input.Index == -1 {
		if input.NamedPosition == protos.NamedPosition_POSITION_TOP {
			list = list.InsertTop(input.Media)
		} else if input.NamedPosition == protos.NamedPosition_POSITION_BOTTOM {
			list = list.InsertBottom(input.Media)
		} else {
			output.Error = "invalid named position and missing index, what should we do"
			return
		}
	} else {
		if list, err = list.InsertAt(input.Media, int(input.Index)); err != nil {
			output.Error = err.Error()
			return
		}
	}
	if err = m.queueSaveToDB(list); err != nil {
		log.Println(err)
		return
	}
	output.Rows = mediasToQueues(list)

	return
}

// handlerQueueRemove removes one media from the queue
func (m *MsgHandler) handlerQueueRemove(msg *nats.Msg) {
	var input protos.QueueInput
	var err error
	var output protos.QueueOutput

	defer func() {
		onQueueNotifyChanged(m.base.NATS())
		if msg.Reply == "" {
			return
		}
		data, err := helpers.ProtoMarshal(&output)
		if err != nil {
			log.Println(err)
			return
		}
		if err = msg.Respond(data); err != nil {
			log.Println(err)
		}
	}()
	if err := helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		output.Error = err.Error()
		log.Println(err)
		return
	}
	ctx := context.Background()
	conn := m.db
	list, err := queueList(ctx, conn)
	if err != nil {
		output.Error = err.Error()
		return
	}
	if list, err = list.RemoveAt(int(input.Index)); err != nil {
		output.Error = err.Error()
		return
	}
	if err = m.queueSaveToDB(list); err != nil {
		log.Println(err)
		return
	}

	output.Rows = mediasToQueues(list)

	return
}

// queueSaveToDB saves queue to db within a db transaction scope
func (m *MsgHandler) queueSaveToDB(medias types.Medias) error {
	conn, err := m.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		if err != nil {
			err = conn.Rollback()
		} else {
			err = conn.Commit()
		}
		if err != nil {
			log.Println(err)
		}
	}()
	ctx := context.Background()
	if err = queueSave(ctx, conn, medias); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
