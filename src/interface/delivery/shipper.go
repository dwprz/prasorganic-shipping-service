package delivery

import (
	"context"

	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
)

type Shipper interface {
	GetProvinces(ctx context.Context) ([]*entity.Province, error)
}
