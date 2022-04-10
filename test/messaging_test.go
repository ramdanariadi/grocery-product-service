package test

import (
	"fmt"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/streadway/amqp"
	"testing"
)

func Test_open_rmq_connection(t *testing.T) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helpers.PanicIfError(err)

	channel, err := connection.Channel()
	helpers.PanicIfError(err)

	queue, err := channel.QueueDeclare(
		"hello2",
		false,
		false,
		false,
		false,
		nil)
	helpers.PanicIfError(err)

	body := "hello world"
	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	helpers.PanicIfError(err)

	defer channel.Close()
	defer connection.Close()
}

func Test_consume_message(t *testing.T) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	helpers.PanicIfError(err)

	channel, err := connection.Channel()
	helpers.PanicIfError(err)

	queue, err := channel.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil)
	helpers.PanicIfError(err)

	msgs, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	helpers.PanicIfError(err)

	go func() {
		for msg := range msgs {
			fmt.Println(msg)
		}
	}()

	defer channel.Close()
	defer connection.Close()
}
