package client

import (
	"encoding/json"
	"errors"
	"fmt"
	amqprpc "github.com/0x4b53/amqp-rpc"
	"github.com/Gewinum/go-weather-microservices/common"
	"github.com/Gewinum/go-weather-microservices/restapi-server/internal/config"
	"github.com/go-viper/mapstructure/v2"
)

func BuildRabbitMQURL(c config.RabbitMQConfig) string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
	)
}

func SendCommand[REQ any, RES any](connection *amqprpc.Client, routingKey string, command REQ) (*RES, error) {
	dataBytes, err := json.Marshal(command)
	if err != nil {
		panic(err)
		return nil, err
	}

	madeRequest := amqprpc.NewRequest().WithRoutingKey(routingKey).WithContentType("application/json").WithBody(string(dataBytes))
	delivery, err := connection.Send(madeRequest)
	if err != nil {
		return nil, err
	}
	var payload common.ResultPayload
	err = json.Unmarshal(delivery.Body, &payload)
	if err != nil {
		return nil, err
	}
	if payload.Error != "" {
		return nil, errors.New(payload.Error)
	}
	var result RES
	err = mapstructure.Decode(payload.Result, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
