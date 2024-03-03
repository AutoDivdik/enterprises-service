//go:build wireinject
// +build wireinject

package app

import (
	"github.com/AutoDivdik/enterprises-service/internal/app/router"
	"github.com/AutoDivdik/enterprises-service/internal/infras/repo"
	"github.com/AutoDivdik/enterprises-service/internal/usecases"
	engine "github.com/AutoDivdik/pg-engine"
	"github.com/AutoDivdik/rmq"
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

func InitApp(dbConnStr engine.DBConnString, connStr rmq.RabbitMQConnStr, grpcServer *grpc.Server) (*App, func(), error) {
	panic(wire.Build(
		NewApplication,
		initDB,
		initRabbitMQ,
		wire.NewSet(rmq.NewConsumer),
		wire.NewSet(router.NewGRPCServer),
		wire.NewSet(repo.NewEnterpriseRepo),
		wire.NewSet(usecases.NewUseCase),
	))
}

func initDB(url engine.DBConnString) (engine.DBEngine, func(), error) {
	db, err := engine.NewPostgresDB(url, 5, 2)
	if err != nil {
		return nil, nil, err
	}

	return db, func() { db.Close() }, nil
}

func initRabbitMQ(url rmq.RabbitMQConnStr) (*amqp.Connection, func(), error) {
	conn, err := rmq.NewRabbitMQConn(url, 5, 2)
	if err != nil {
		return nil, nil, err
	}

	return conn, func() { conn.Close() }, nil
}
