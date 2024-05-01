package queue

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Sniff(queueName string, hostname string, port int, username string, password string) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", username, password, hostname, port))
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to broker: %s", err.Error()))
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to channel: %s", err.Error()))
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		panic(fmt.Sprintf("Unable to consume message: %s", err.Error()))
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
		}
	}()

	<-forever
}
