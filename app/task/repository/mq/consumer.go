package mq

import (
	"context"

	"github.com/streadway/amqp"
)

func ConnsumeMessage(ctx context.Context, queueName string) (msg <-chan amqp.Delivery, err error) {
	ch, err := Rabbitmq.Channel()
	if err != nil {
		panic(err)
	}

	q, _ := ch.QueueDeclare(queueName, true, false, false, false, nil)
	err = ch.Qos(1, 0, false)
	if err != nil {
		panic(err)
	}
	return ch.Consume(q.Name, "", false, false, false, false, nil)
}
