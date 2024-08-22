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
// go test -run ^TestSetvice_GetSuburbsByCityId$ -v ./src/service/test/ -count=1

type GetSuburbsByCityIdTestSuite struct {
	suite.Suite
	shippingService service.Shipping
	shipperRestful  *delivery.ShipperMock
	shippingCache   *cache.ShippingMock
}

func (g *GetSuburbsByCityIdTestSuite) SetupSuite() {
	orderGrpc := delivery.NewOrderGrpcMock()
	orderConn := new(grpc.ClientConn)

	grpcClient := grpcclient.NewGrpc(orderGrpc, orderConn)
	g.shippingCache = cache.NewShippingMock()

	g.shipperRestful = delivery.NewShipperMock()
	restfulClient := restfulclient.NewRestful(g.shipperRestful)

	g.shippingService = serviceimpl.NewShipping(restfulClient, grpcClient, g.shippingCache)
}

func (g *GetSuburbsByCityIdTestSuite) Test_Success() {
	cityId := 6
	shipperRes := g.CreateShipperRes()

	g.shippingCache.Mock.On("FindSuburbsByCityId", mock.Anything, cityId).Return(nil)
	g.shipperRestful.Mock.On("GetSuburbsByCityId", mock.Anything, cityId).Return(shipperRes, nil)

	res, err := g.shippingService.GetSuburbsByCityId(context.Background(), cityId)
	assert.NoError(g.T(), err)

	assert.Equal(g.T(), shipperRes, res)
}

func (g *GetSuburbsByCityIdTestSuite) Test_GetFromCache() {
	cityId := 7
	shipperRes := g.CreateShipperRes()

	g.shippingCache.Mock.On("FindSuburbsByCityId", mock.Anything, cityId).Return(shipperRes)

	res, err := g.shippingService.GetSuburbsByCityId(context.Background(), cityId)
	assert.NoError(g.T(), err)

	assert.Equal(g.T(), shipperRes, res)
}

func (g *GetSuburbsByCityIdTestSuite) CreateShipperRes() *dto.ShipperRes[[]*entity.Suburb] {
	suburbs := g.CreateSuburbs()

	return &dto.ShipperRes[[]*entity.Suburb]{
		Data: suburbs,
		Pagination: &entity.Pagination{
			CurrentPage:   1,
			TotalPages:    1,
			TotalElements: 6,
		},
	}
}

func (g *GetSuburbsByCityIdTestSuite) CreateSuburbs() []*entity.Suburb {
	return []*entity.Suburb{
		{
			Id:       482,
			Name:     "Setia Budi",
			Lat:      -6.2195686,
			Lng:      106.8325872,
			City:     &entity.City{Id: 41, Name: "Jakarta Selatan", Lat: -6.2614927, Lng: 106.8105998},
			Province: &entity.Province{Id: 6, Name: "DKI Jakarta", Lat: -6.1744651, Lng: 106.822745},
			Country:  &entity.Country{Id: 228, Name: "INDONESIA", Code: "ID"},
		},
		{
			Id:       483,
			Name:     "Tebet",
			Lat:      -6.2318597,
			Lng:      106.8473377,
			City:     &entity.City{Id: 41, Name: "Jakarta Selatan", Lat: -6.2614927, Lng: 106.8105998},
			Province: &entity.Province{Id: 6, Name: "DKI Jakarta", Lat: -6.1744651, Lng: 106.822745},
			Country:  &entity.Country{Id: 228, Name: "INDONESIA", Code: "ID"},
		},
		{
			Id:       484,
			Name:     "Mampang Prapatan",
			Lat:      -6.2506144,
			Lng:      106.8207875,
			City:     &entity.City{Id: 41, Name: "Jakarta Selatan", Lat: -6.2614927, Lng: 106.8105998},
			Province: &entity.Province{Id: 6, Name: "DKI Jakarta", Lat: -6.1744651, Lng: 106.822745},
			Country:  &entity.Country{Id: 228, Name: "INDONESIA", Code: "ID"},
		},
		{
			Id:       485,
			Name:     "Pancoran",
			Lat:      -6.2523005,
			Lng:      106.8473377,
			City:     &entity.City{Id: 41, Name: "Jakarta Selatan", Lat: -6.2614927, Lng: 106.8105998},
			Province: &entity.Province{Id: 6, Name: "DKI Jakarta", Lat: -6.1744651, Lng: 106.822745},
			Country:  &entity.Country{Id: 228, Name: "INDONESIA", Code: "ID"},
		},
		{
			Id:       486,
			Name:     "Jagakarsa",
			Lat:      -6.334917,
			Lng:      106.8237374,
			City:     &entity.City{Id: 41, Name: "Jakarta Selatan", Lat: -6.2614927, Lng: 106.8105998},
			Province: &entity.Province{Id: 6, Name: "DKI Jakarta", Lat: -6.1744651, Lng: 106.822745},
			Country:  &entity.Country{Id: 228, Name: "INDONESIA", Code: "ID"},
		},
		{
			Id:       487,
			Name:     "Pasar Minggu",
			Lat:      -6.2939813,
			Lng:      106.8237374,
			City:     &entity.City{Id: 41, Name: "Jakarta Selatan", Lat: -6.2614927, Lng: 106.8105998},
			Province: &entity.Province{Id: 6, Name: "DKI Jakarta", Lat: -6.1744651, Lng: 106.822745},
			Country:  &entity.Country{Id: 228, Name: "INDONESIA", Code: "ID"},
		},
		{
			Id:       488,
			Name:     "Cilandak",
			Lat:      -6.2845276,
			Lng:      106.8001396,
			City:     &entity.City{Id: 41, Name: "Jakarta Selatan", Lat: -6.2614927, Lng: 106.8105998},
			Province: &entity.Province{Id: 6, Name: "DKI Jakarta", Lat: -6.1744651, Lng: 106.822745},
			Country:  &entity.Country{Id: 228, Name: "INDONESIA", Code: "ID"},
		},
		{
			Id:       489,
			Name:     "Pesanggrahan",
			Lat:      -6.2474283,
			Lng:      106.7617984,
			City:     &entity.City{Id: 41, Name: "Jakarta Selatan", Lat: -6.2614927, Lng: 106.8105998},
			Province: &entity.Province{Id: 6, Name: "DKI Jakarta", Lat: -6.1744651, Lng: 106.822745},
			Country:  &entity.Country{Id: 228, Name: "INDONESIA", Code: "ID"},
		},
		{
			Id:       490,
			Name:     "Kebayoran Lama",
			Lat:      -6.2443916,
			Lng:      106.7765443,
			City:     &entity.City{Id: 41, Name: "Jakarta Selatan", Lat: -6.2614927, Lng: 106.8105998},
			Province: &entity.Province{Id: 6, Name: "DKI Jakarta", Lat: -6.1744651, Lng: 106.822745},
			Country:  &entity.Country{Id: 228, Name: "INDONESIA", Code: "ID"},
		},
		{
			Id:       491,
			Name:     "Kebayoran Baru",
			Lat:      -6.2436219,
			Lng:      106.8001396,
			City:     &entity.City{Id: 41, Name: "Jakarta Selatan", Lat: -6.2614927, Lng: 106.8105998},
			Province: &entity.Province{Id: 6, Name: "DKI Jakarta", Lat: -6.1744651, Lng: 106.822745},
			Country:  &entity.Country{Id: 228, Name: "INDONESIA", Code: "ID"},
		},
	}
}
func TestSetvice_GetSuburbsByCityId(t *testing.T) {
	suite.Run(t, new(GetSuburbsByCityIdTestSuite))
}
