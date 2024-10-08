package main

import rabbitmq "github.com/garciawell/go-full-pos/29-eventos/events/pkg/rabbitMQ"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello, World!", "amq.direct")
}
