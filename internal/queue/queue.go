package queue

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Sniff(queueName string, hostname string, port int, username string, password string) {
	// TODO: Is there a way to figure out which messaging broker is being used based on the provided connection
	// information? This would prevent us from having to have the user specify that information.

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
		nil,       // args
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
