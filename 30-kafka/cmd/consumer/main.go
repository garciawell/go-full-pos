package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "30-kafka-kafka-1:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group",
	}

	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		panic(err)
	}

	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println("Mensagem recebida: ", string(msg.Value), msg.TopicPartition)
		}
	}
}
