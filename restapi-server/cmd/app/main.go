package main

import (
	"context"
	"github.com/Gewinum/go-weather-microservices/restapi-server/internal/dependency"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	interruption, cancel1 := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel1()
	dpd := dependency.NewDependency()
	go func() {
		err := dpd.GinEngine.Run(net.JoinHostPort(dpd.Config.Server.Host, dpd.Config.Server.Port))
		if err != nil {
			panic(err)
		}
	}()
	<-interruption.Done()
	log.Println("Shutting down")
}
