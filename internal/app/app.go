package app

import (
	"context"
	"github.com/AutoDivdik/enterprises-service/gen"
	"github.com/AutoDivdik/enterprises-service/internal/usecases"
	engine "github.com/AutoDivdik/pg-engine"
	"github.com/AutoDivdik/rmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type App struct {
	AMQPConn *amqp.Connection
	PG       engine.DBEngine

	Consumer rmq.EventConsumer

	UC                   usecases.UseCase
	EnterpriseGRPCServer gen.EnterpriseServiceServer
}

func NewApplication(
	pg engine.DBEngine,
	amqpConn *amqp.Connection,

	consumer rmq.EventConsumer,
	uc usecases.UseCase,
	enterpriseGRPCServer gen.EnterpriseServiceServer,
) *App {
	return &App{
		AMQPConn: amqpConn,
		PG:       pg,

		Consumer: consumer,

		UC:                   uc,
		EnterpriseGRPCServer: enterpriseGRPCServer,
	}
}

func (c *App) Worker(ctx context.Context, messages <-chan amqp.Delivery) {

}
