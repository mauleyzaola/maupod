package main

import (
	"context"
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/conversion"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/types"
	"github.com/nats-io/nats.go"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func queueList(ctx context.Context, conn boil.ContextExecutor) (types.Medias, error) {
	rows, err := orm.MediaQueues(qm.OrderBy(orm.MediaQueueColumns.ID+" asc")).All(ctx, conn)
	if err != nil {
		return nil, err
	}

	var ids []interface{}
	for _, v := range rows {
		ids = append(ids, v.MediaID)
	}

	medias, err := orm.Media(qm.WhereIn(orm.MediumColumns.ID+" in ?", ids...)).All(ctx, conn)
	if err != nil {
		return nil, err
	}
	var keys = make(map[string]*pb.Media)
	for _, v := range medias {
		keys[v.ID] = conversion.MediaFromORM(v)
	}

	var res []*pb.Media
	for _, v := range rows {
		val, ok := keys[v.MediaID]
		if !ok {
			continue
		}
		res = append(res, val)
	}
	return res, nil
}

// TODO: improve this shit
func queueSave(ctx context.Context, conn boil.ContextExecutor, rows types.Medias) error {
	_, err := orm.MediaQueues().DeleteAll(ctx, conn)
	if err != nil {
		return err
	}
	for i, v := range rows {
		q := orm.MediaQueue{
			ID:       helpers.NewUUID(),
			Position: i,
			MediaID:  v.Id,
		}
		if err = q.Insert(ctx, conn, boil.Infer()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MsgHandler) handlerQueueList(msg *nats.Msg) {
	var output pb.QueueOutput

	defer func() {
		data, err := helpers.ProtoMarshal(&output)
		if err != nil {
			log.Println(err)
			return
		}
		if err = msg.Respond(data); err != nil {
			log.Println(err)
		}
	}()
	output.Rows = mediasToQueues(m.queueItems)
}

func (m *MsgHandler) handlerQueueAdd(msg *nats.Msg) {
	var input pb.QueueInput
	var err error
	var output pb.QueueOutput

	defer func() {
		if err = m.queueSave(); err != nil {
			log.Println(err)
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

	list := m.queueItems

	if input.Index == -1 {
		if input.NamedPosition == pb.NamedPosition_POSITION_TOP {
			list = list.InsertTop(input.Media)
		} else if input.NamedPosition == pb.NamedPosition_POSITION_BOTTOM {
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
	m.queueItems = list
	output.Rows = mediasToQueues(list)
	return
}

func (m *MsgHandler) handlerQueueRemove(msg *nats.Msg) {
	var input pb.QueueInput
	var err error
	var output pb.QueueOutput

	defer func() {
		if err = m.queueSave(); err != nil {
			log.Println(err)
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
	list := m.queueItems
	if list, err = list.RemoveAt(int(input.Index)); err != nil {
		return
	}
	m.queueItems = list
	output.Rows = mediasToQueues(list)
}

func mediasToQueues(medias types.Medias) []*pb.Queue {
	var result []*pb.Queue
	for i, v := range medias {
		result = append(result, &pb.Queue{
			Id:       helpers.NewUUID(),
			Media:    v,
			Position: int32(i),
		})
	}
	return result
}

func (m *MsgHandler) queueSave() error {
	conn, err := m.db.Begin()
	if err != nil {
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
	if _, err = conn.Exec("truncate table " + orm.TableNames.MediaQueue); err != nil {
		return err
	}
	log.Printf("[INFO] persisting %d queue items to db\n", len(m.queueItems))
	ctx := context.Background()
	for i, v := range m.queueItems {
		item := &orm.MediaQueue{
			ID:       helpers.NewUUID(),
			Position: i,
			MediaID:  v.Id,
		}
		if err = item.Insert(ctx, conn, boil.Infer()); err != nil {
			return err
		}
	}
	return err
}

func (m *MsgHandler) handlerQueueSave(msg *nats.Msg) {
	if err := m.queueSave(); err != nil {
		log.Println(err)
	}
}
