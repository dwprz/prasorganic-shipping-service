package cache

import (
	"context"

	"github.com/dwprz/prasorganic-shipping-service/src/model/dto"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
)

type Shipping interface {
	CacheProvinces(ctx context.Context, p *dto.ShipperRes[[]*entity.Province])
	CacheCitiesByProvinceId(ctx context.Context, provinceId int, p *dto.ShipperRes[[]*entity.City])
	FindProvinces(ctx context.Context) *dto.ShipperRes[[]*entity.Province]
	FindCitiesByProvinceId(ctx context.Context, provinceId int) *dto.ShipperRes[[]*entity.City]
}
