package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	fmt.Println("Kafka Producer")
	producer := NewKafkaProducer()
	err := Publish("mensagem 1", "teste", producer, nil)
	if err != nil {
		log.Println("EPAA")
	}
	log.Println("Producer criado")
	producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "30-kafka-kafka-1:9092",
	}

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println("Error creating producer", err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte) error {

	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}

	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}

	return nil
}
