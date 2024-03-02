package app

import (
	"context"
	"github.com/AutoDivdik/rmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type App struct {
	AMQPConn *amqp.Connection

	Consumer rmq.EventConsumer
}

func NewApplication(amqpConn *amqp.Connection, consumer rmq.EventConsumer) *App {
	return &App{
		AMQPConn: amqpConn,
		Consumer: consumer,
	}
}

func (c *App) Worker(ctx context.Context, messages <-chan amqp.Delivery) {

}
