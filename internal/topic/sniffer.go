package topic

import (
	"fmt"

	"github.com/joshmalbrecht/message-sniffer/internal/connection"
)

func Sniff(exchangeName string, bindingKey string, hostname string, port int, virtHost string, username string, password string) {
	conn := connection.Connect(hostname, port, virtHost, username, password)

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to channel: %s", err.Error()))
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		true,  // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		panic(fmt.Sprintf("Unable to declare queue: %s", err.Error()))
	}

	err = ch.QueueBind(
		q.Name,       // queue name
		bindingKey,   // key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		panic(fmt.Sprintf("Unable to bind queue: %s", err.Error()))
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // arguments
	)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
		}
	}()

	<-forever
}
