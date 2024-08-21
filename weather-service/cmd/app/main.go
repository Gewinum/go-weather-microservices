package main

import (
	"context"
	"github.com/Gewinum/go-weather-microservices/weather-service/internal/dependency"
	"os"
	"os/signal"
)

func main() {
	interruption, cancel1 := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel1()
	dpd := dependency.NewDependency()
	go func() {
		dpd.RpcServer.ListenAndServe()
	}()
	<-interruption.Done()
}
