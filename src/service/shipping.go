package service

import (
	"context"

	"github.com/dwprz/prasorganic-shipping-service/src/core/restful/client"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/cache"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/service"
	"github.com/dwprz/prasorganic-shipping-service/src/model/dto"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
)

type ShippingImpl struct {
	restfulClient *client.Restful
	shippingCache cache.Shipping
}

func NewShipping(rc *client.Restful, sc cache.Shipping) service.Shipping {
	return &ShippingImpl{
		restfulClient: rc,
		shippingCache: sc,
	}
}

func (s *ShippingImpl) GetProvinces(ctx context.Context) (*dto.ShipperRes[[]*entity.Province], error) {
	if res := s.shippingCache.FindProvinces(ctx); res != nil {
		return res, nil
	}

	res, err := s.restfulClient.Shipper.GetProvinces(ctx)
	if err == nil && len(res.Data) > 0 {
		go s.shippingCache.CacheProvinces(context.Background(), res)
	}

	return res, err
}

func (s *ShippingImpl) GetCitiesByProvinceId(ctx context.Context, provinceId int) (*dto.ShipperRes[[]*entity.City], error) {
	if res := s.shippingCache.FindCitiesByProvinceId(ctx, provinceId); res != nil {
		return res, nil
	}

	res, err := s.restfulClient.Shipper.GetCitiesByProvinceId(ctx, provinceId)
	if err == nil && len(res.Data) > 0 {
		go s.shippingCache.CacheCitiesByProvinceId(context.Background(), provinceId, res)
	}

	return res, err
}

func (s *ShippingImpl) GetSuburbsByCityId(ctx context.Context, cityId int) (*dto.ShipperRes[[]*entity.Suburb], error) {
	if res := s.shippingCache.FindSuburbsByCityId(ctx, cityId); res != nil {
		return res, nil
	}

	res, err := s.restfulClient.Shipper.GetSuburbsByCityId(ctx, cityId)
	if err == nil && len(res.Data) > 0 {
		go s.shippingCache.CacheSuburbsByCityId(context.Background(), cityId, res)
	}

	return res, err
}
