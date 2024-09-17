package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"golang.org/x/exp/rand"
)

func main() {
	producer := NewKafkaProducer()
	rand.Seed(uint64(time.Now().UnixNano()))
	randomNumber := rand.Intn(100)

	deliveryCh := make(chan kafka.Event)

	Publish(fmt.Sprintf("Transferiu %d", randomNumber), "teste", producer, []byte("transferencia2"), deliveryCh)

	go DeliveryReport(deliveryCh)
	producer.Flush(5000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "30-kafka-kafka-1:9092",
		"delivery.timeout.ms": "0",
		"acks":                "all",
		"enable.idempotence":  "true",
	}

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println("Error creating producer", err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryCh chan kafka.Event) error {

	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}

	err := producer.Produce(message, deliveryCh)
	if err != nil {
		return err
	}

	return nil
}

func DeliveryReport(deliveryChannel chan kafka.Event) {
	for e := range deliveryChannel {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar a mensagem", ev.TopicPartition.Error)
			} else {
				fmt.Println("Mensagem entregue", ev.TopicPartition)
			}
		}
	}
}
