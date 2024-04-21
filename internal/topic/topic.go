package topic

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func sniff(hostname string, port int, username string, password string) {
	// TODO: Is there a way to figure out which messaging broker is being used based on the provided connection
	// information? This would prevent us from having to have the user specify that information.

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", username, password, hostname, password))
	if err != nil {
		// TODO: handle
	}

	defer conn.Close()

}
