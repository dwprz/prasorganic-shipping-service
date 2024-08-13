package service

import (
	"context"

	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
)

type Shipping interface {
	GetProvinces(ctx context.Context) ([]*entity.Province, error)
}
