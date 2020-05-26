package handler

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
	"github.com/nats-io/nats.go"
)

type MsgHandler struct {
	logger types.Logger
	nc     *nats.Conn

	subscriptions []*nats.Subscription
}

type Subscription struct {
	Subject string
	Handler nats.MsgHandler
}

func NewMsgHandler(logger types.Logger, nc *nats.Conn) *MsgHandler {
	return &MsgHandler{
		logger: logger,
		nc:     nc,
	}
}

func (m *MsgHandler) Register(subs ...Subscription) error {
	var err error
	var sub *nats.Subscription
	for _, s := range subs {
		if sub, err = m.nc.Subscribe(s.Subject, s.Handler); err != nil {
			m.logger.Error(err)
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
			m.logger.Error(err)
		}
	}
	m.nc.Close()
}

func (m *MsgHandler) Logger() types.Logger {
	return m.logger
}
