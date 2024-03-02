package main

import (
	"context"
	"enterprises/internal/app"
	"fmt"
	"github.com/AutoDivdik/rmq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	a, cleanup, err := app.InitApp(rmq.RabbitMQConnStr("amqp://rmquser:rmqpassword@localhost:5672/"))
	if err != nil {
		cancel()
		log.Fatalf("failed init app: %v", err)
	}

	go func() {
		err := a.Consumer.Consume(a.Worker)
		if err != nil {
			cancel()
			log.Fatalf("failed to start consumer: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	log.Println("start enterprise service...")

	select {
	case v := <-quit:
		cleanup()
		fmt.Printf("signal.notify %v", v)
	case done := <-ctx.Done():
		cleanup()
		fmt.Printf("ctx.done %v", done)
	}
}
