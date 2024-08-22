package test

import (
	"context"
	"testing"

	grpcclient "github.com/dwprz/prasorganic-shipping-service/src/core/grpc/client"
	restfulclient "github.com/dwprz/prasorganic-shipping-service/src/core/restful/client"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/service"
	"github.com/dwprz/prasorganic-shipping-service/src/mock/cache"
	"github.com/dwprz/prasorganic-shipping-service/src/mock/delivery"
	"github.com/dwprz/prasorganic-shipping-service/src/model/dto"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
	serviceimpl "github.com/dwprz/prasorganic-shipping-service/src/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

// go test -p=1 -v ./src/service/test/... -count=1
// go test -run ^TestSetvice_GetCitiesProvinceId$ -v ./src/service/test/ -count=1

type GetCitiesByProvinceIdTestSuite struct {
	suite.Suite
	shippingService service.Shipping
	shipperRestful  *delivery.ShipperMock
	shippingCache   *cache.ShippingMock
}

func (g *GetCitiesByProvinceIdTestSuite) SetupSuite() {
	orderGrpc := delivery.NewOrderGrpcMock()
	orderConn := new(grpc.ClientConn)

	grpcClient := grpcclient.NewGrpc(orderGrpc, orderConn)
	g.shippingCache = cache.NewShippingMock()

	g.shipperRestful = delivery.NewShipperMock()
	restfulClient := restfulclient.NewRestful(g.shipperRestful)

	g.shippingService = serviceimpl.NewShipping(restfulClient, grpcClient, g.shippingCache)
}

func (g *GetCitiesByProvinceIdTestSuite) Test_Success() {
	provinceId := 6
	shipperRes := g.CreateShipperRes()

	g.shippingCache.Mock.On("FindCitiesByProvinceId", mock.Anything, provinceId).Return(nil)
	g.shipperRestful.Mock.On("GetCitiesByProvinceId", mock.Anything, provinceId).Return(shipperRes, nil)

	res, err := g.shippingService.GetCitiesByProvinceId(context.Background(), provinceId)
	assert.NoError(g.T(), err)

	assert.Equal(g.T(), shipperRes, res)
}

func (g *GetCitiesByProvinceIdTestSuite) Test_GetFromCache() {
	provinceId := 7
	shipperRes := g.CreateShipperRes()

	g.shippingCache.Mock.On("FindCitiesByProvinceId", mock.Anything, provinceId).Return(shipperRes)

	res, err := g.shippingService.GetCitiesByProvinceId(context.Background(), provinceId)
	assert.NoError(g.T(), err)

	assert.Equal(g.T(), shipperRes, res)
}

func (g *GetCitiesByProvinceIdTestSuite) CreateShipperRes() *dto.ShipperRes[[]*entity.City] {
	cities := g.CreateCities()

	return &dto.ShipperRes[[]*entity.City]{
		Data: cities,
		Pagination: &entity.Pagination{
			CurrentPage:   1,
			TotalPages:    1,
			TotalElements: 6,
		},
	}
}

func (g *GetCitiesByProvinceIdTestSuite) CreateCities() []*entity.City {
	return []*entity.City{
		{
			Province: &entity.Province{
				Id:   6,
				Name: "DKI Jakarta",
				Lat:  -6.1744651,
				Lng:  106.822745,
			},
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
			Id:   38,
			Name: "Kepulauan Seribu",
			Lat:  -5.7985265,
			Lng:  106.5071981,
		},
		{
			Province: &entity.Province{
				Id:   6,
				Name: "DKI Jakarta",
				Lat:  -6.1744651,
				Lng:  106.822745,
			},
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
			Id:   39,
			Name: "Jakarta Utara",
			Lat:  -6.1384145,
			Lng:  106.863956,
		},
		{
			Province: &entity.Province{
				Id:   6,
				Name: "DKI Jakarta",
				Lat:  -6.1744651,
				Lng:  106.822745,
			},
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
			Id:   40,
			Name: "Jakarta Timur",
			Lat:  -6.2250138,
			Lng:  106.9004472,
		},
		{
			Province: &entity.Province{
				Id:   6,
				Name: "DKI Jakarta",
				Lat:  -6.1744651,
				Lng:  106.822745,
			},
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
			Id:   41,
			Name: "Jakarta Selatan",
			Lat:  -6.2614927,
			Lng:  106.8105998,
		},
		{
			Province: &entity.Province{
				Id:   6,
				Name: "DKI Jakarta",
				Lat:  -6.1744651,
				Lng:  106.822745,
			},
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
			Id:   42,
			Name: "Jakarta Barat",
			Lat:  -6.1683295,
			Lng:  106.7588494,
		},
		{
			Province: &entity.Province{
				Id:   6,
				Name: "DKI Jakarta",
				Lat:  -6.1744651,
				Lng:  106.822745,
			},
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
			Id:   43,
			Name: "Jakarta Pusat",
			Lat:  -6.1864864,
			Lng:  106.8340911,
		},
	}
}
func TestSetvice_GetCitiesProvinceId(t *testing.T) {
	suite.Run(t, new(GetCitiesByProvinceIdTestSuite))
}
