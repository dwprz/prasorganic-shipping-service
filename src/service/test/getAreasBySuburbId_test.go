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
// go test -run ^TestSetvice_GetAreasBySuburbId$ -v ./src/service/test/ -count=1

type GetAreasBySuburbIdTestSuite struct {
	suite.Suite
	shippingService service.Shipping
	shipperRestful  *delivery.ShipperMock
	shippingCache   *cache.ShippingMock
}

func (g *GetAreasBySuburbIdTestSuite) SetupSuite() {
	orderGrpc := delivery.NewOrderGrpcMock()
	orderConn := new(grpc.ClientConn)

	grpcClient := grpcclient.NewGrpc(orderGrpc, orderConn)
	g.shippingCache = cache.NewShippingMock()

	g.shipperRestful = delivery.NewShipperMock()
	restfulClient := restfulclient.NewRestful(g.shipperRestful)

	g.shippingService = serviceimpl.NewShipping(restfulClient, grpcClient, g.shippingCache)
}

func (g *GetAreasBySuburbIdTestSuite) Test_Success() {
	suburbId := 6
	shipperRes := g.CreateShipperRes()

	g.shippingCache.Mock.On("FindAreasBySuburbId", mock.Anything, suburbId).Return(nil)
	g.shipperRestful.Mock.On("GetAreasBySuburbId", mock.Anything, suburbId).Return(shipperRes, nil)

	res, err := g.shippingService.GetAreasBySuburbId(context.Background(), suburbId)
	assert.NoError(g.T(), err)

	assert.Equal(g.T(), shipperRes, res)
}

func (g *GetAreasBySuburbIdTestSuite) Test_GetFromCache() {
	suburbId := 7
	shipperRes := g.CreateShipperRes()

	g.shippingCache.Mock.On("FindAreasBySuburbId", mock.Anything, suburbId).Return(shipperRes)

	res, err := g.shippingService.GetAreasBySuburbId(context.Background(), suburbId)
	assert.NoError(g.T(), err)

	assert.Equal(g.T(), shipperRes, res)
}

func (g *GetAreasBySuburbIdTestSuite) CreateShipperRes() *dto.ShipperRes[[]*entity.Area] {
	areas := g.CreateAreas()

	return &dto.ShipperRes[[]*entity.Area]{
		Data: areas,
		Pagination: &entity.Pagination{
			CurrentPage:   1,
			TotalPages:    1,
			TotalElements: 8,
		},
	}
}

func (g *GetAreasBySuburbIdTestSuite) CreateAreas() []*entity.Area {
	return []*entity.Area{
		{
            Suburb: &entity.Suburb{
                Id:   482,
                Name: "Setia Budi",
                Lat:  -6.2195686,
                Lng:  106.8325872,
            },
            City: &entity.City{
                Id:   41,
                Name: "Jakarta Selatan",
                Lat:  -6.2614927,
                Lng:  106.8105998,
            },
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
            Id:       4707,
            Name:     "Pasar Manggis",
            Postcode: "12970",
            Lat:      -6.2090339,
            Lng:      106.8415828,
        },
        {
            Suburb: &entity.Suburb{
                Id:   482,
                Name: "Setia Budi",
                Lat:  -6.2195686,
                Lng:  106.8325872,
            },
            City: &entity.City{
                Id:   41,
                Name: "Jakarta Selatan",
                Lat:  -6.2614927,
                Lng:  106.8105998,
            },
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
            Id:       4708,
            Name:     "Guntur",
            Postcode: "12980",
            Lat:      -6.2080273,
            Lng:      106.8340622,
        },
        {
            Suburb: &entity.Suburb{
                Id:   482,
                Name: "Setia Budi",
                Lat:  -6.2195686,
                Lng:  106.8325872,
            },
            City: &entity.City{
                Id:   41,
                Name: "Jakarta Selatan",
                Lat:  -6.2614927,
                Lng:  106.8105998,
            },
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
            Id:       4709,
            Name:     "Kuningan Timur",
            Postcode: "12950",
            Lat:      -6.2299792,
            Lng:      106.8266873,
        },
        {
            Suburb: &entity.Suburb{
                Id:   482,
                Name: "Setia Budi",
                Lat:  -6.2195686,
                Lng:  106.8325872,
            },
            City: &entity.City{
                Id:   41,
                Name: "Jakarta Selatan",
                Lat:  -6.2614927,
                Lng:  106.8105998,
            },
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
            Id:       4710,
            Name:     "Menteng Atas",
            Postcode: "12960",
            Lat:      -6.2180511,
            Lng:      106.8399623,
        },
        {
            Suburb: &entity.Suburb{
                Id:   482,
                Name: "Setia Budi",
                Lat:  -6.2195686,
                Lng:  106.8325872,
            },
            City: &entity.City{
                Id:   41,
                Name: "Jakarta Selatan",
                Lat:  -6.2614927,
                Lng:  106.8105998,
            },
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
            Id:       4711,
            Name:     "Karet Kuningan",
            Postcode: "12940",
            Lat:      -6.2197608,
            Lng:      106.8266873,
        },
        {
            Suburb: &entity.Suburb{
                Id:   482,
                Name: "Setia Budi",
                Lat:  -6.2195686,
                Lng:  106.8325872,
            },
            City: &entity.City{
                Id:   41,
                Name: "Jakarta Selatan",
                Lat:  -6.2614927,
                Lng:  106.8105998,
            },
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
            Id:       4712,
            Name:     "Karet Semanggi",
            Postcode: "12930",
            Lat:      -6.2213742,
            Lng:      106.8163628,
        },
        {
            Suburb: &entity.Suburb{
                Id:   482,
                Name: "Setia Budi",
                Lat:  -6.2195686,
                Lng:  106.8325872,
            },
            City: &entity.City{
                Id:   41,
                Name: "Jakarta Selatan",
                Lat:  -6.2614927,
                Lng:  106.8105998,
            },
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
            Id:       4713,
            Name:     "Setiabudi",
            Postcode: "12910",
            Lat:      -6.2098892,
            Lng:      106.8223457,
        },
        {
            Suburb: &entity.Suburb{
                Id:   482,
                Name: "Setia Budi",
                Lat:  -6.2195686,
                Lng:  106.8325872,
            },
            City: &entity.City{
                Id:   41,
                Name: "Jakarta Selatan",
                Lat:  -6.2614927,
                Lng:  106.8105998,
            },
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
            Id:       4714,
            Name:     "Karet",
            Postcode: "12920",
            Lat:      -6.214731,
            Lng:      106.818265,
        },

	}
}
func TestSetvice_GetAreasBySuburbId(t *testing.T) {
	suite.Run(t, new(GetAreasBySuburbIdTestSuite))
}
