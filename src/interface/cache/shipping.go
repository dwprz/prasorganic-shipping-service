package cache

import (
	"context"

	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
)

type Shipping interface {
	CacheProvinces(ctx context.Context, p []*entity.Province)
	FindProvinces(ctx context.Context) []*entity.Province
}
