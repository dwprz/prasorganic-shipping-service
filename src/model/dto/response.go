package dto

type ShipperResponse[T any] struct {
	Data T `json:"data"`
}
