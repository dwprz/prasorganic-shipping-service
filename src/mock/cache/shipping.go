package cache

import (
	"context"

	"github.com/dwprz/prasorganic-shipping-service/src/model/dto"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
	"github.com/stretchr/testify/mock"
)

type ShippingMock struct {
	mock.Mock
}

func NewShippingMock() *ShippingMock {
	return &ShippingMock{
		Mock: mock.Mock{},
	}
}

func (s *ShippingMock) CacheTrackingByShippingId(ctx context.Context, shippingId string, t *entity.Tracking) {
}

func (s *ShippingMock) CacheProvinces(ctx context.Context, p *dto.ShipperRes[[]*entity.Province]) {}

func (s *ShippingMock) CacheCitiesByProvinceId(ctx context.Context, provinceId int, p *dto.ShipperRes[[]*entity.City]) {
}

func (s *ShippingMock) CacheSuburbsByCityId(ctx context.Context, cityId int, p *dto.ShipperRes[[]*entity.Suburb]) {
}

func (s *ShippingMock) CacheAreasBySuburbId(ctx context.Context, suburbId int, p *dto.ShipperRes[[]*entity.Area]) {
}

func (s *ShippingMock) FindTrackingByShippingId(ctx context.Context, shippingId string) *entity.Tracking {
	arguments := s.Mock.Called(ctx, shippingId)

	if arguments.Get(0) == nil {
		return nil
	}

	return arguments.Get(0).(*entity.Tracking)
}

func (s *ShippingMock) FindProvinces(ctx context.Context) *dto.ShipperRes[[]*entity.Province] {
	arguments := s.Mock.Called(ctx)

	if arguments.Get(0) == nil {
		return nil
	}

	return arguments.Get(0).(*dto.ShipperRes[[]*entity.Province])
}

func (s *ShippingMock) FindCitiesByProvinceId(ctx context.Context, provinceId int) *dto.ShipperRes[[]*entity.City] {
	arguments := s.Mock.Called(ctx, provinceId)

	if arguments.Get(0) == nil {
		return nil
	}

	return arguments.Get(0).(*dto.ShipperRes[[]*entity.City])
}

func (s *ShippingMock) FindSuburbsByCityId(ctx context.Context, cityId int) *dto.ShipperRes[[]*entity.Suburb] {
	arguments := s.Mock.Called(ctx, cityId)

	if arguments.Get(0) == nil {
		return nil
	}

	return arguments.Get(0).(*dto.ShipperRes[[]*entity.Suburb])
}

func (s *ShippingMock) FindAreasBySuburbId(ctx context.Context, suburbId int) *dto.ShipperRes[[]*entity.Area] {
	arguments := s.Mock.Called(ctx, suburbId)

	if arguments.Get(0) == nil {
		return nil
	}

	return arguments.Get(0).(*dto.ShipperRes[[]*entity.Area])
}

func (s *ShippingMock) UpdateTracking(ctx context.Context, shippingId string, data *entity.TrackingData) {
}
