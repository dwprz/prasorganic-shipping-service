package dto

import "github.com/dwprz/prasorganic-shipping-service/src/model/entity"

type ShipperRes[T any] struct {
	Data       T `json:"data"`
	Pagination *entity.Pagination `json:"pagination,omitempty"`
}
