package main

import "github.com/nats-io/nats.go"

func (m *MsgHandler) handlerPing(msg *nats.Msg) {
	if err := msg.Respond(nil); err != nil {
		m.base.Logger().Error(err)
	}
}
