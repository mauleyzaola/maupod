package handler

import (
	"log"

	"github.com/nats-io/nats.go"
)

type MsgHandler struct {
	nc *nats.Conn

	subscriptions []*nats.Subscription
}

type Subscription struct {
	Subject string
	Handler nats.MsgHandler
}

func NewMsgHandler(nc *nats.Conn) *MsgHandler {
	return &MsgHandler{
		nc: nc,
	}
}

func (m *MsgHandler) Register(subs ...Subscription) error {
	var err error
	var sub *nats.Subscription
	for _, s := range subs {
		if sub, err = m.nc.Subscribe(s.Subject, s.Handler); err != nil {
			log.Println(err)
			return err
		}
		m.subscriptions = append(m.subscriptions, sub)
	}
	return nil
}

func (m *MsgHandler) Close() {
	var err error
	for _, sub := range m.subscriptions {
		if err = sub.Unsubscribe(); err != nil {
			log.Println(err)
		}
	}
	m.nc.Close()
}

func (m *MsgHandler) NATS() *nats.Conn {
	return m.nc
}
