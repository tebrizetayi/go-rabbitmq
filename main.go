package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ Tutorial")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	fmt.Println("Successfully Connected to our rabbitmq instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("TestQueue", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish("", "TestQueue", false, false,
		amqp.Publishing{ContentType: "text/plain", Body: []byte("Hello World " + time.Now().String())})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully Published message to queue")
}
