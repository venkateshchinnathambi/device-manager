package handler

import (
	thermostatpb "device-manager/proto"
)

func init() {
	Register(&thermostatpb.IndoorTemperatueUpdated{}, &IndoorTemperatureHandler{})
}
