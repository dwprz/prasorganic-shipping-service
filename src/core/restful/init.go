package restful

import (
	"github.com/dwprz/prasorganic-shipping-service/src/core/restful/client"
	"github.com/dwprz/prasorganic-shipping-service/src/core/restful/delivery"
	"github.com/dwprz/prasorganic-shipping-service/src/core/restful/handler"
	"github.com/dwprz/prasorganic-shipping-service/src/core/restful/middleware"
	"github.com/dwprz/prasorganic-shipping-service/src/core/restful/server"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/service"
)

func InitServer(ss service.Shipping) *server.Restful {
	shippingHandler := handler.NewShipping(ss)
	middleware := middleware.New()

	restfulServer := server.NewRestful(shippingHandler, middleware)
	return restfulServer
}

func InitClient() *client.Restful {
	shipperDelivery := delivery.NewShipper()
	restfulClient := client.NewRestful(shipperDelivery)

	return restfulClient
}
