package broker

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var Channel *amqp.Channel

func Connect() *amqp.Channel {
	conn, err := amqp.Dial("amqp://" + "myuser" + ":" + "mypassword" + "@" + "172.17.0.5" + ":" + "5672" + "/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	// Use the connection to create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	Channel = ch
	fmt.Print("rabbit connected 2")
	return ch
}
func GetBrokerConnection() *amqp.Channel {
	return Channel
}
