package service

import (
	"context"

	"github.com/dwprz/prasorganic-shipping-service/src/model/dto"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
)

type Shipping interface {
	GetProvinces(ctx context.Context) (*dto.ShipperRes[[]*entity.Province], error)
	GetCitiesByProvinceId(ctx context.Context, provinceId int) (*dto.ShipperRes[[]*entity.City], error)
}
