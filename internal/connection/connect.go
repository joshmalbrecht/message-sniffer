package connection

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect(hostname string, port int, virtHost string, username string, password string) *amqp.Connection {
	if len(virtHost) != 0 {
		virtHost = "/" + virtHost
	}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d%s", username, password, hostname, port, virtHost))
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to broker: %s", err.Error()))
	}

	return conn
}
