package mq

import (
	"fmt"
	"micro-todolist/config"

	"github.com/streadway/amqp"
)

var Rabbitmq *amqp.Connection

func InitRabbitMQ() {
	connString := fmt.Sprintf("%s://%s:%s@%s:%s/", config.RabbitMQ, config.RabbitMQUser, config.RabbitMQPassWord, config.RabbitMQHost, config.RabbitMQPort)
	fmt.Println("InitRabbitMQ...", connString)
	conn, err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	Rabbitmq = conn
}
