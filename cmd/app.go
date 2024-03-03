package main

import (
	"context"
	"fmt"
	"github.com/AutoDivdik/enterprises-service/internal/app"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	srv := grpc.NewServer()
	go func() {
		defer srv.GracefulStop()
		<-ctx.Done()
	}()

	a, cleanup, err := app.InitApp("postgres://root:root@localhost:5432/enterprises?sslmode=disable", "amqp://rmquser:rmqpassword@localhost:5672/", srv)
	if err != nil {
		cancel()
		log.Fatalf("failed init app: %v", err)
	}

	l, err := net.Listen("tcp", ":2020")
	if err != nil {
		cancel()
		<-ctx.Done()
		log.Fatalf("failed to listen to address :2020")
	}

	log.Println("start enterprise service...")
	defer func() {
		if err1 := l.Close(); err != nil {
			<-ctx.Done()
			log.Fatalf("%v", err1)
		}
	}()
	go func() {
		err := a.Consumer.Consume(a.Worker)
		if err != nil {
			cancel()
			log.Fatalf("failed to start consumer: %v", err)
		}
	}()

	if err = srv.Serve(l); err != nil {
		cancel()
		<-ctx.Done()
		log.Fatalf("failed to start gRPC server")
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		cleanup()
		fmt.Printf("signal.notify %v", v)
	case done := <-ctx.Done():
		cleanup()
		fmt.Printf("ctx.done %v", done)
	}
}
