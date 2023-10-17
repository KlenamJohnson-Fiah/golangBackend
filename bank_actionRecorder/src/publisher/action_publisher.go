package publisher

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// type RabbitMqInterface interface {
// }

// type rabbitMQ struct {
// 	conn *amqp.Connection
// }

// func NewRabbiteMqConnection(conn *amqp.Connection) RabbitMqInterface {
// 	return &rabbitMQ{
// 		conn: conn,
// 	}
// }

var conn *amqp.Connection

func DBInit(c string) {
	var err error

	conn, err := amqp.Dial(c)
	if err != nil {
		log.Fatalf("could not connect to rabbitmq: %v", err)
		panic(err)
	}
	conn.Close()
}

func Publish(q string, msg []byte) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	payload := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         msg,
	}

	if err := ch.Publish("", q, false, false, payload); err != nil {
		return fmt.Errorf("[Publish] failed to publish to queue %v", err)
	}

	return nil
}
