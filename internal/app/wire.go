//go:build wireinject
// +build wireinject

package app

import (
	"github.com/AutoDivdik/rmq"
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitApp(connStr rmq.RabbitMQConnStr) (*App, func(), error) {
	panic(wire.Build(
		NewApplication, initRabbitMQ, wire.NewSet(rmq.NewConsumer),
	))
}

func initRabbitMQ(url rmq.RabbitMQConnStr) (*amqp.Connection, func(), error) {
	conn, err := rmq.NewRabbitMQConn(url, 5, 2)
	if err != nil {
		return nil, nil, err
	}

	return conn, func() { conn.Close() }, nil
}
