package topic

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Sniff(exchangeName string, routingKey string, hostname string, port int, username string, password string) {
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
		routingKey,   // routing key
		exchangeName, // exchange
		false,
		nil,
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
		nil,    // args
	)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
		}
	}()

	<-forever
}
