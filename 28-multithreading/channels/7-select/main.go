package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Msg string
	id  int64
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{Msg: "Hello RabbitMQ", id: i}
			c1 <- msg
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second)
			msg := Message{Msg: "Hello Kafka", id: i}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1: //rabbitMq
			fmt.Printf("Received from RabbitMQ: ID %d -  %s\n", msg.id, msg.Msg)
		case msg := <-c2: //rabit kafka
			fmt.Printf("Received from Kafka: ID %d -  %s\n", msg.id, msg.Msg)

		case <-time.After(time.Second * 3):
			println("timeout")

			// default:
			// 	println("default")
		}
	}

}
