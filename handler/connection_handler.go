package handler

import (
	thermostatpb "device-manager/proto"
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
		var event thermostatpb.Event
		if err := proto.Unmarshal(buf[:n], &event); err != nil {
			log.Printf("Error Parsing Proto %v", err)
			continue
		}
		dispatchEvent(&event)
	}
}
