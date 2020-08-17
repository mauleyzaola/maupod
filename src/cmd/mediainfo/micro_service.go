package main

import (
	"log"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerMicroService(msg *nats.Msg) {
	if err := broker.MicroServiceRespond(msg, helpers.AppName()); err != nil {
		log.Println(err)
	}
}
