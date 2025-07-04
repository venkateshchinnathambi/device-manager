package main

import (
<<<<<<< HEAD
	thermostatpb "device-manager/proto"
	"fmt"
	"net"

=======
	"fmt"
	"net"

	thermostatpb "device-manager/proto"

>>>>>>> 134670a8013b8ff2a4ed8d17bb538af3c78ae4ee
	"google.golang.org/protobuf/proto"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:7778")
	if err != nil {
		fmt.Printf("Unable to connect %v", err)
		return
	}

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
	data, err := proto.Marshal(event)
	if err != nil {
		fmt.Printf("Error encode proto %v", err)
		return
	}
	conn.Write(data)
}
