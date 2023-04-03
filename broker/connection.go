package broker

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var Channel *amqp.Channel

func Connect() *amqp.Channel {
	conn, err := amqp.Dial("amqp://" + "myuser" + ":" + "mypassword" + "@" + "10.105.66.16" + ":" + "5672" + "/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	fmt.Print("rabbit connected")
	// Use the connection to create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	Channel = ch
	return ch
}
func GetBrokerConnection() *amqp.Channel {
	return Channel
}
