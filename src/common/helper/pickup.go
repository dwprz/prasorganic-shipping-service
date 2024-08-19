package helper

import "github.com/dwprz/prasorganic-shipping-service/src/model/entity"

func FormatPickupReq(shippingIds []string) *entity.Pickup {
	return &entity.Pickup{
		Data: struct{OrderActivation entity.OrderActivation "json:\"order_activation\""}{
			entity.OrderActivation{
				OrderId: shippingIds,
			},
		},
	}
}
