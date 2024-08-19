package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dwprz/prasorganic-shipping-service/src/cache"
	"github.com/dwprz/prasorganic-shipping-service/src/core/broker"
	"github.com/dwprz/prasorganic-shipping-service/src/core/grpc"
	"github.com/dwprz/prasorganic-shipping-service/src/core/restful"
	"github.com/dwprz/prasorganic-shipping-service/src/infrastructure/database"
	"github.com/dwprz/prasorganic-shipping-service/src/service"
)

func handleCloseApp(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	handleCloseApp(cancel)

	redisDB := database.NewRedisCluster()
	shippingCache := cache.NewShipping(redisDB)

	grpcClient := grpc.InitClient()

	restfulClient := restful.InitClient()
	shippingService := service.NewShipping(restfulClient, grpcClient, shippingCache)

	restfulServer := restful.InitServer(shippingService)
	defer restfulServer.Stop()

	go restfulServer.Run()

	notifService := service.NewNotification(grpcClient, shippingCache)

	shipperConsumer := broker.InitShipperConsumer(notifService)
	defer shipperConsumer.Close()

	go shipperConsumer.Consume(ctx)

	<-ctx.Done()
}
