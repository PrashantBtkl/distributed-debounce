package main

import (
	"github.com/PrashantBtkl/distributed-debounce/debouncer/api"
	log "github.com/sirupsen/logrus"
	"time"
)

// func run() {
// 	log.SetHandler(cli.Default)
// 	log.SetLevel(log.DebugLevel)

// 	var config model.Config

// 	l, err := api.NewHTTPListener("127.0.0.1", "1234")
// 	if err != nil {
// 		log.Fatalf("run: failed to init config: %v", err)
// 	}

// 	rmq, err := rabbitmq.InitRabbitMQ(config.AMQP)
// 	if err != nil {
// 		log.Fatalf("run: failed to init rabbitmq: %v", err)
// 	}
// 	defer rmq.Shutdown()

// 	err = rmq.PublishWithDelay("user.event.publish", []byte("hello"), 10000)
// 	if err != nil {
// 		log.Fatalf("run: failed to publish into rabbitmq: %v", err)
// 	}

// 	for {
// 	}
// }

func httpListener(done chan<- int) {
	defer func() {
		done <- 1
	}()

	retries := 0
	for retries < 5 {
		err := api.NewHTTPListener("0.0.0.0", "1234")
		if err != nil {
			log.WithFields(log.Fields{
				"err": err.Error(),
			}).Error("Failed to create Listener: Exiting")
			retries++
			time.Sleep(3 * time.Second)
			continue
		}
		time.Sleep(3 * time.Second)
		retries++
	}
	log.Error("Retries Over: Exiting")
}

func main() {
	done := make(chan int)
	go httpListener(done)
	<-done
}
