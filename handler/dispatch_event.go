package handler

import (
	thermostatpb "device-manager/proto"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func dispatchEvent(event *thermostatpb.Event) error {
	msg, ok := getOneOfMessge(event)
	if ok != nil {
		return fmt.Errorf("invalid event received %v", ok)
	}
	fullName := string(msg.ProtoReflect().Descriptor().FullName())
	handle := Registery[fullName]

	return handle.Handle(msg)
}

func getOneOfMessge(event *thermostatpb.Event) (proto.Message, error) {
	if event == nil {
		return nil, fmt.Errorf("invalid event")
	}

	oneof := event.GetEventType()
	switch e := oneof.(type) {
	case *thermostatpb.Event_IndoorTemperatureUpdated:
		return e.IndoorTemperatureUpdated, nil
	case *thermostatpb.Event_SetPointsUpdated:
		return e.SetPointsUpdated, nil
	default:
		return nil, fmt.Errorf("invalid event type")
	}
}
