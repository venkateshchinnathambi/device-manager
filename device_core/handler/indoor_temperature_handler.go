package handler

import (
	"device-manager/models"
	thermostatpb "device-manager/proto"
	"device-manager/utils"
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
)

type IndoorTemperatureHandler struct{}

type IndoorTemperature struct {
	ZoneId      string
	ZoneName    string
	Temperature string
}

func (h *IndoorTemperatureHandler) Handle(msg proto.Message) error {
	idt, ok := msg.(*thermostatpb.IndoorTemperatueUpdated)
	if !ok {
		return fmt.Errorf("invalid message type %v", ok)
	}
	fmt.Printf("Message Received from: %v - IndoorTemperatueUpdated - Value: %v\n",
		idt.GetDeviceId().GetId(),
		idt.GetTemperature().GetValue(),
	)
	utils.DB.Create(&models.IndoorTemperature{
		DeviceID:  idt.GetDeviceId().GetId(),
		Value:     idt.GetTemperature().GetValue(),
		Scale:     idt.GetTemperature().GetScale().String(),
		Timestamp: int64(idt.GetTimestamp()),
	})

	jsonMsg := IndoorTemperature{
		ZoneId:      "1",
		ZoneName:    "Living Room",
		Temperature: idt.GetTemperature().GetValue(),
	}
	msgBytes, err := json.Marshal(jsonMsg)
	if err != nil {
		return fmt.Errorf("encoding error %v", err)
	}
	utils.SendKafkaMessage("devices.updates", string(msgBytes))
	return nil
}
