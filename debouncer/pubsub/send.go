package mq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func testAMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	args := make(amqp.Table)
	args["x-delayed-type"] = "direct"
	err = ch.ExchangeDeclare("delayed", "x-delayed-message", true, false, false, false, args)
	if err != nil {
		log.Fatal("Failed declaring exchange")
	}

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(q.Name, "user.event.publish", "delayed", false, nil)
	failOnError(err, "failed to bind queue with exchange")

	headers := make(amqp.Table)
	headers["x-delay"] = 5000

	body := "Hello World!"
	err = ch.Publish(
		"delayed", // exchange
		q.Name,    // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
			Headers:      headers,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)

	err = ch.Qos(2, 0, false)
	if err != nil {
		log.Println(err)
	}

	msgs, err := ch.Consume(
		q.Name,    // queue
		"consume", // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")

	for {
		log.Println("start listenning to consume")

		select {
		case d, ok := <-msgs:
			if !ok {
				return
			}
			log.Println("consume : ", string(d.Body))
			d.Ack(false)
		}
	}
}
