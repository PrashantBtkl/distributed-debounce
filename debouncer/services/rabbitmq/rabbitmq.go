package rabbitmq

import (
	"github.com/PrashantBtkl/distributed-debounce/debouncer/model"
)

type rmq interface {
	InitRabbitMQ(config model.AMQP) (*model.RabbitMQ, error)
	Publish(routingKey string, body []byte) error
	PublishWithDelay(routingKey string, body []byte, delay int64) error
	Shutdown()
}
