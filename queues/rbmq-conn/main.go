package main

import (
	"fmt"
	"internal"
)

type App struct {
	Rmq *internal.RabbitMQ
}

func Run() error {
	fmt.Println("RabbitMQ Runner.")

	rbmq := internal.NewRBMQService()
	app := App{
		Rmq: rbmq,
	}

	err := app.Rmq.Connect()
	if err != nil {
		return err
	}
	defer app.Rmq.Conn.Close()

	err = app.Rmq.Publish("HI!")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Error setting up the consumer....")
		fmt.Println(err)
	}
}