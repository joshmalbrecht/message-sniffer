package queue

import (
	"fmt"

	"github.com/joshmalbrecht/message-sniffer/internal/connection"
)

func Sniff(queueName string, hostname string, port int, virtHost string, username string, password string) {
	conn := connection.Connect(hostname, port, virtHost, username, password)

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
