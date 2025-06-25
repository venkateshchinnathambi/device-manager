package main

import (
	"device-manager/handler"
	"device-manager/models"
	"device-manager/utils"
	"fmt"
	"log"
	"net"
)

func main() {
	utils.InitDB()
	utils.DB.AutoMigrate(&models.IndoorTemperature{})
	utils.InitKafkaProducer()
	listner, err := net.Listen("tcp", "0.0.0.0:7777")
	if err != nil {
		log.Fatalf("Error option TCP on 7777")
		return
	}
	defer listner.Close()
	fmt.Println("Server Listening on 7777")

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatalf("Error Accepting Connections %v", err)
			continue
		}
		go handler.HandleConnection(conn)
	}

}
