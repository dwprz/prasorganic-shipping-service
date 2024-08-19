package helper

import "github.com/dwprz/prasorganic-shipping-service/src/model/entity"

func FormatPickupReq(shippingIds []string) *entity.Pickup {
	return &entity.Pickup{
		Data: entity.PickupData{
			OrderActivation: entity.OrderActivation{
				OrderId: shippingIds,
			},
		},
	}
}
