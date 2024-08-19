package entity

type Pickup struct {
	Data Data `json:"data"`
}

type Data struct {
	OrderActivation OrderActivation `json:"order_activation"`
}

type OrderActivation struct {
	OrderId []string `json:"order_id"`
}
