package entity

type Pickup struct {
	Data PickupData `json:"data"`
}

type PickupData struct {
	OrderActivation OrderActivation `json:"order_activation"`
}

type OrderActivation struct {
	OrderId []string `json:"order_id"`
}
