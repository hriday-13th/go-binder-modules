package internal

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Service
type Service interface {
	Connect() error
	Publish(message string) error
}

// RabbitMQ
type RabbitMQ struct {
	Conn 		*amqp.Connection
	Channel 	*amqp.Channel
}

// Connect - establihes a connection to your RBMQ instance
// and declares the queue
func (r *RabbitMQ) Connect() error {
	fmt.Println("Connecting to rbmq...")

	var err error
	r.Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to rbmq.")

	r.Channel, err = r.Conn.Channel()
	if err != nil {
		return err
	}

	_, err = r.Channel.QueueDeclare(
		"TEST-QUEUE",
		false,
		false,
		false,
		false,
		nil,
	)

	return nil
}

// Publish - publishes a message to the queue
func (r *RabbitMQ) Publish(message string) error {
	err := r.Channel.Publish(
		"",
		"TEST-QUEUE",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		},
	)

	if err != nil {
		return err
	}

	fmt.Println("Successfully published message to queue.")
	return nil
}

func NewRBMQService() *RabbitMQ {
	return &RabbitMQ{}
}