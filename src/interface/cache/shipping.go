package cache

import (
	"context"

	"github.com/dwprz/prasorganic-shipping-service/src/model/dto"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
)

type Shipping interface {
	CacheTrackingByShippingId(ctx context.Context, shippingId string, t *entity.Tracking)
	CacheProvinces(ctx context.Context, p *dto.ShipperRes[[]*entity.Province])
	CacheCitiesByProvinceId(ctx context.Context, provinceId int, p *dto.ShipperRes[[]*entity.City])
	CacheSuburbsByCityId(ctx context.Context, cityId int, p *dto.ShipperRes[[]*entity.Suburb])
	CacheAreasBySuburbId(ctx context.Context, suburbId int, p *dto.ShipperRes[[]*entity.Area])

	FindTrackingByShippingId(ctx context.Context, shippingId string) *entity.Tracking
	FindProvinces(ctx context.Context) *dto.ShipperRes[[]*entity.Province]
	FindCitiesByProvinceId(ctx context.Context, provinceId int) *dto.ShipperRes[[]*entity.City]
	FindSuburbsByCityId(ctx context.Context, cityId int) *dto.ShipperRes[[]*entity.Suburb]
	FindAreasBySuburbId(ctx context.Context, suburbId int) *dto.ShipperRes[[]*entity.Area]
	UpdateTracking(ctx context.Context, shippingId string, data *entity.TrackingData)
}
