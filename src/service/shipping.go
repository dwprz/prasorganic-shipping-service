package service

import (
	"context"

	pb "github.com/dwprz/prasorganic-proto/protogen/order"
	grpcclient "github.com/dwprz/prasorganic-shipping-service/src/core/grpc/client"
	resfulclient "github.com/dwprz/prasorganic-shipping-service/src/core/restful/client"
	v "github.com/dwprz/prasorganic-shipping-service/src/infrastructure/validator"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/cache"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/service"
	"github.com/dwprz/prasorganic-shipping-service/src/model/dto"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
)

type ShippingImpl struct {
	restfulClient *resfulclient.Restful
	grpcClient    *grpcclient.Grpc
	shippingCache cache.Shipping
}

func NewShipping(rc *resfulclient.Restful, gc *grpcclient.Grpc, sc cache.Shipping) service.Shipping {
	return &ShippingImpl{
		restfulClient: rc,
		grpcClient:    gc,
		shippingCache: sc,
	}
}

func (s *ShippingImpl) ShippingOrder(ctx context.Context, data *entity.ShippingOrder) error {
	if err := v.Validate.Struct(data); err != nil {
		return err
	}

	orderId := data.ExternalId

	shippingId, err := s.restfulClient.Shipper.ShippingOrder(ctx, data)
	if err != nil {
		return err
	}

	err = s.grpcClient.Order.AddShippingId(ctx, &pb.AddShippingIdReq{
		OrderId:    orderId,
		ShippingId: shippingId,
	})

	if err != nil {
		return err
	}

	err = s.restfulClient.Shipper.RequestPickup(ctx, []string{shippingId})

	return err
}

func (s *ShippingImpl) Pricing(ctx context.Context, data *dto.PricingReq) (*dto.ShipperRes[*entity.Pricing], error) {
	if err := v.Validate.Struct(data); err != nil {
		return nil, err
	}

	res, err := s.restfulClient.Shipper.Pricing(ctx, data)
	return res, err
}

func (s *ShippingImpl) CreateLabel(ctx context.Context, data *dto.CreateLabelReq) (*dto.ShipperRes[*entity.Label], error) {
	if err := v.Validate.Struct(data); err != nil {
		return nil, err
	}

	res, err := s.restfulClient.Shipper.CreateLabel(ctx, data)
	return res, err
}

func (s *ShippingImpl) TrackingByShippingId(ctx context.Context, shippingId string) (*entity.Tracking, error) {
	if res := s.shippingCache.FindTrackingByShippingId(ctx, shippingId); res != nil {
		return res, nil
	}

	res, err := s.restfulClient.Shipper.TrackingByShippingId(ctx, shippingId)
	if err == nil && len(res.Trackings) > 0 {
		go s.shippingCache.CacheTrackingByShippingId(context.Background(), shippingId, res)
	}

	return res, err
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

func (s *ShippingImpl) GetAreasBySuburbId(ctx context.Context, suburbId int) (*dto.ShipperRes[[]*entity.Area], error) {
	if res := s.shippingCache.FindAreasBySuburbId(ctx, suburbId); res != nil {
		return res, nil
	}

	res, err := s.restfulClient.Shipper.GetAreasBySuburbId(ctx, suburbId)
	if err == nil && len(res.Data) > 0 {
		go s.shippingCache.CacheAreasBySuburbId(context.Background(), suburbId, res)
	}

	return res, err
}
