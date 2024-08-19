package delivery

import (
	"context"

	"github.com/dwprz/prasorganic-shipping-service/src/model/dto"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
)

type Shipper interface {
	ShippingOrder(ctx context.Context, data *entity.ShippingOrder) (shippingId string, err error)
	RequestPickup(ctx context.Context, shippingIds []string) error
	Pricing(ctx context.Context, data *dto.PricingReq) (*dto.ShipperRes[*entity.Pricing], error)
	GetProvinces(ctx context.Context) (*dto.ShipperRes[[]*entity.Province], error)
	GetCitiesByProvinceId(ctx context.Context, provinceId int) (*dto.ShipperRes[[]*entity.City], error)
	GetSuburbsByCityId(ctx context.Context, cityId int) (*dto.ShipperRes[[]*entity.Suburb], error)
	GetAreasBySuburbId(ctx context.Context, suburbId int) (*dto.ShipperRes[[]*entity.Area], error)
}
