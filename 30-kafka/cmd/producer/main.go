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

	Publish(fmt.Sprintf("Mensagem %d", randomNumber), "teste", producer, nil, deliveryCh)

	e := <-deliveryCh
	msg := e.(*kafka.Message)

	if msg.TopicPartition.Error != nil {
		fmt.Println("Erro ao enviar a mensagem", msg.TopicPartition.Error)
		return
	} else {
		fmt.Println("Mensagem entregue", msg.TopicPartition)
	}

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
