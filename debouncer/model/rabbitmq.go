package model

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	URL        string
	Exchange   string
	Conn       *amqp.Connection
	Chann      *amqp.Channel
	Queue      amqp.Queue
	CloseChann chan *amqp.Error
	QuitChann  chan bool
}
