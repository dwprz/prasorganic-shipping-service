package client

import "github.com/dwprz/prasorganic-shipping-service/src/interface/delivery"

type Restful struct {
	Shipper delivery.Shipper
}

func NewRestful(sd delivery.Shipper) *Restful {
	return &Restful{
		Shipper: sd,
	}
}
