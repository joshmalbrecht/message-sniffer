package queue

import (
	"fmt"

	"github.com/joshmalbrecht/message-sniffer/internal/connection"
)

func Sniff(queueName string, hostname string, port int, virtHost string, username string, password string) error {
	conn := connection.Connect(hostname, port, virtHost, username, password)

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(fmt.Sprintf("Unable to connect to channel: %s", err.Error()))
		return err
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		queueName, // queue name
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		fmt.Println(fmt.Sprintf("Unable to consume message: %s", err.Error()))
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
		}
	}()

	<-forever

	return nil
}
