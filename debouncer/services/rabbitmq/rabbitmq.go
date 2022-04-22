package rabbitmq

import (
	"github.com/PrashantBtkl/distributed-debounce/debouncer/model"
	"github.com/PrashantBtkl/distributed-debounce/debouncer/services/store"
)

type rmq interface {
	InitRabbitMQ(config model.AMQP, db *store.PGStore) (*model.RabbitMQ, error)
	Publish(routingKey string, body []byte) error
	PublishWithDelay(routingKey string, body []byte, delay int64) error
	Shutdown()
}
