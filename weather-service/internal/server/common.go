package server

import (
	"context"
	"encoding/json"
	"fmt"
	amqprpc "github.com/0x4b53/amqp-rpc"
	"github.com/Gewinum/go-weather-microservices/common"
	"github.com/Gewinum/go-weather-microservices/weather-service/internal/config"
	"github.com/streadway/amqp"
)

type HandlerFunc[REQ any, RES any] func(command REQ) (RES, error)

func NewRPCServer(c config.RabbitMQConfig) *amqprpc.Server {
	return amqprpc.NewServer(BuildRabbitMQURL(c))
}

func BuildRabbitMQURL(c config.RabbitMQConfig) string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
	)
}

func BindHandler[REQ any, RES any](connection *amqprpc.Server, routingKey string, handler HandlerFunc[REQ, RES]) {
	connection.Bind(amqprpc.DirectBinding(routingKey, func(ctx context.Context, writer *amqprpc.ResponseWriter, delivery amqp.Delivery) {
		var cmdInstance REQ
		err := json.Unmarshal(delivery.Body, &cmdInstance)
		if err != nil {
			_, _ = writer.Write(ErrorPayload(err))
			return
		}
		result, err := handler(cmdInstance)
		if err != nil {
			_, _ = writer.Write(ErrorPayload(err))
			return
		}
		if err != nil {
			_, _ = writer.Write(ErrorPayload(err))
			return
		}
		_, _ = writer.Write(SuccessPayload(result))
	}))
}

func SuccessPayload(data any) []byte {
	payload := common.ResultPayload{
		Error:  "",
		Result: data,
	}
	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	return bytesPayload
}

func ErrorPayload(err error) []byte {
	payload := common.ResultPayload{
		Error:  err.Error(),
		Result: nil,
	}
	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	return bytesPayload
}
