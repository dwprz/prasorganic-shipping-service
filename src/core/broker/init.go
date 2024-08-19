package broker

import (
	"github.com/dwprz/prasorganic-shipping-service/src/core/broker/consumer"
	"github.com/dwprz/prasorganic-shipping-service/src/core/broker/handler"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/service"
)

func InitShipperConsumer(ns service.Notification) *consumer.ShipperKafka {
	shipperHandler := handler.NewShipperKafka(ns)
	shipperConsumer := consumer.NewShipperKafka(shipperHandler)

	return shipperConsumer
}
