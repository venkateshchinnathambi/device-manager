package utils

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var producer *ckafka.Producer

func InitKafkaProducer() {
	var err error
	producer, err = ckafka.NewProducer(&ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092", // Docker service name
	})

	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *ckafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()
}

func SendKafkaMessage(topic string, value string) {
	err := producer.Produce(&ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(value),
	}, nil)

	if err != nil {
		log.Printf("Error producing message: %v\n", err)
	}
}
