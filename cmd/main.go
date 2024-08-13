package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dwprz/prasorganic-shipping-service/src/cache"
	"github.com/dwprz/prasorganic-shipping-service/src/core/restful"
	"github.com/dwprz/prasorganic-shipping-service/src/infrastructure/database"
	"github.com/dwprz/prasorganic-shipping-service/src/service"
)

func handleCloseApp(closeCH chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		close(closeCH)
	}()
}

func main() {
	closeCH := make(chan struct{})
	handleCloseApp(closeCH)

	redisDB := database.NewRedisCluster()
	shippingCache := cache.NewShipping(redisDB)

	restfulClient := restful.InitClient()
	shippingService := service.NewShipping(restfulClient, shippingCache)

	restfulServer := restful.InitServer(shippingService)
	defer restfulServer.Stop()

	go restfulServer.Run()

	<-closeCH
}
