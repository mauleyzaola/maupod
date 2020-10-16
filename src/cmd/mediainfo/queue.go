package main

import (
	"context"
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/conversion"
	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/types"
	"github.com/nats-io/nats.go"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// queueList returns a list of media sorted by the queue position
func queueList(ctx context.Context, conn boil.ContextExecutor) (types.Medias, error) {
	rows, err := orm.MediaQueues(qm.OrderBy(orm.MediaQueueColumns.Position+" asc")).All(ctx, conn)
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

// queueSave drops all items from queue and inserts them again in the provided order
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

// onQueueNotifyChanged notifies listeners that queue has been changed
func onQueueNotifyChanged(nc *nats.Conn) {
	var input pb.QueueChangedInput
	if err := broker.PublishBrokerJSON(nc, pb.Message_MESSAGE_SOCKET_QUEUE_CHANGE, &input); err != nil {
		log.Println(err)
		return
	}
}

// mediasToQueues converts medias to queues, taking the position from the slice index
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
