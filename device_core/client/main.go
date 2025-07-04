package main

import (
	thermostatpb "device-manager/proto"
	"fmt"
	"net"

	"google.golang.org/protobuf/proto"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:7778")
	if err != nil {
		fmt.Printf("Unable to connect %v", err)
		return
	}
	//eyJhbGciOiJSUzI1NiJ9.eyJkZXZpY2VfaWQiOiIxMjM0IiwiZXhwIjoxNzUxNjMyOTIyfQ.NUEBa2GgZPxAqvrD2M77DoZxuZkY8PBBCaeIspwEboqDrqfrvAtUvqZRvsDQyoHcR_-fHd4Z24K2enS2h6EfkHabO4WMDdG576UeVyiFoB9ls2FgePK-HNd6D4DwST2rUIke4aSdlJsHKLbW-6gG0tKUZLNqRAdDsd4Is1E9LYz00WZtD0TfHBIlTJ2o7Xz1kxzLqTR_pGSvgtIsYHoOe4eUbaBG94qa2NZmWRVSnk60X4XZvhJ0GKaW758YVrA1Tz97qWZGO7LlTWcIURZAAWCoO7_audupxYcHa7ux5swyHx73q0r3Sz9cng0TtJwFLfkd35XMdLdMWCiU8rgywQ
	event := &thermostatpb.Event{
		EventType: &thermostatpb.Event_IndoorTemperatureUpdated{
			IndoorTemperatureUpdated: &thermostatpb.IndoorTemperatueUpdated{
				DeviceId: &thermostatpb.DeviceId{
					Id: proto.String("1234"),
				},
				Temperature: &thermostatpb.Temperature{
					Value: proto.String("70"),
					Scale: thermostatpb.Temperature_F.Enum(),
				},
				Timestamp: proto.Int32(124567868),
			},
		},
	}
	deviceMsg := &thermostatpb.DeviceMessage{
		Token: "eyJhbGciOiJSUzI1NiJ9.eyJkZXZpY2VfaWQiOiIxMjM0IiwiZXhwIjoxNzUxNjMyOTIyfQ.NUEBa2GgZPxAqvrD2M77DoZxuZkY8PBBCaeIspwEboqDrqfrvAtUvqZRvsDQyoHcR_-fHd4Z24K2enS2h6EfkHabO4WMDdG576UeVyiFoB9ls2FgePK-HNd6D4DwST2rUIke4aSdlJsHKLbW-6gG0tKUZLNqRAdDsd4Is1E9LYz00WZtD0TfHBIlTJ2o7Xz1kxzLqTR_pGSvgtIsYHoOe4eUbaBG94qa2NZmWRVSnk60X4XZvhJ0GKaW758YVrA1Tz97qWZGO7LlTWcIURZAAWCoO7_audupxYcHa7ux5swyHx73q0r3Sz9cng0TtJwFLfkd35XMdLdMWCiU8rgywQ",
		Event: event,
	}
	data, err := proto.Marshal(deviceMsg)
	if err != nil {
		fmt.Printf("Error encode proto %v", err)
		return
	}
	conn.Write(data)
}
