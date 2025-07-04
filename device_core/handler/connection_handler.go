package handler

import (
	thermostatpb "device-manager/proto"
	"device-manager/utils"
	"io"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("Connection Closed ")
			} else {
				log.Printf("Error Reading %v", err)
			}
			return
		}
		var message thermostatpb.DeviceMessage
		if err := proto.Unmarshal(buf[:n], &message); err != nil {
			log.Printf("Error Parsing Proto %v", err)
			continue
		}

		claims, err := utils.ValidateToken(message.GetToken())
		if err != nil {
			log.Println("invalid token:", err)
			return
		}

		log.Printf("Authenticated Device %v", claims.DeviceID)

		var event thermostatpb.Event = *message.GetEvent()
		//	if err := proto.Unmarshal(buf[:n], &event); err != nil {
		//			log.Printf("Error Parsing Proto %v", err)
		//			continue
		//}

		dispatchEvent(&event)
	}
}
