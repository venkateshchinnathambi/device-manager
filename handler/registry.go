package handler

import (
	"google.golang.org/protobuf/proto"
)

type EventHandler interface {
	Handle(proto.Message) error
}

var Registery = map[string]EventHandler{}

func Register(message proto.Message, handler EventHandler) {
	fullName := string(message.ProtoReflect().Descriptor().FullName())
	Registery[fullName] = handler
}
