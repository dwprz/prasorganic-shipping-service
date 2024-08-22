package delivery

import (
	"context"

	"github.com/dwprz/prasorganic-shipping-service/src/model/dto"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
	"github.com/stretchr/testify/mock"
)

type ShipperMock struct {
	mock.Mock
}

func NewShipperMock() *ShipperMock {
	return &ShipperMock{
		Mock: mock.Mock{},
	}
}

func (s *ShipperMock) ShippingOrder(ctx context.Context, data *entity.ShippingOrder) (shippingId string, err error) {
	arguments := s.Mock.Called(ctx, data)

	return arguments.Get(0).(string), arguments.Error(1)
}

func (s *ShipperMock) RequestPickup(ctx context.Context, shippingIds []string) error {
	arguments := s.Mock.Called(ctx, shippingIds)

	return arguments.Error(0)
}

func (s *ShipperMock) Pricing(ctx context.Context, data *dto.PricingReq) (*dto.ShipperRes[*entity.Pricing], error) {
	arguments := s.Mock.Called(ctx, data)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*dto.ShipperRes[*entity.Pricing]), arguments.Error(1)
}

func (s *ShipperMock) CreateLabel(ctx context.Context, data *dto.CreateLabelReq) (*dto.ShipperRes[*entity.Label], error) {
	arguments := s.Mock.Called(ctx, data)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*dto.ShipperRes[*entity.Label]), arguments.Error(1)
}

func (s *ShipperMock) TrackingByShippingId(ctx context.Context, shippingId string) (*entity.Tracking, error) {
	arguments := s.Mock.Called(ctx, shippingId)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*entity.Tracking), arguments.Error(1)
}

func (s *ShipperMock) GetProvinces(ctx context.Context) (*dto.ShipperRes[[]*entity.Province], error) {
	arguments := s.Mock.Called(ctx)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*dto.ShipperRes[[]*entity.Province]), arguments.Error(1)
}

func (s *ShipperMock) GetCitiesByProvinceId(ctx context.Context, provinceId int) (*dto.ShipperRes[[]*entity.City], error) {
	arguments := s.Mock.Called(ctx, provinceId)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*dto.ShipperRes[[]*entity.City]), arguments.Error(1)
}

func (s *ShipperMock) GetSuburbsByCityId(ctx context.Context, cityId int) (*dto.ShipperRes[[]*entity.Suburb], error) {
	arguments := s.Mock.Called(ctx, cityId)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*dto.ShipperRes[[]*entity.Suburb]), arguments.Error(1)
}

func (s *ShipperMock) GetAreasBySuburbId(ctx context.Context, suburbId int) (*dto.ShipperRes[[]*entity.Area], error) {
	arguments := s.Mock.Called(ctx, suburbId)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	return arguments.Get(0).(*dto.ShipperRes[[]*entity.Area]), arguments.Error(1)
}
