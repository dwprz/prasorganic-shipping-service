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
// go test -run ^TestSetvice_GetProvinces$ -v ./src/service/test/ -count=1

type GetProvincesTestSuite struct {
	suite.Suite
	shippingService service.Shipping
	shipperRestful  *delivery.ShipperMock
	shippingCache   *cache.ShippingMock
}

func (g *GetProvincesTestSuite) SetupSuite() {
	orderGrpc := delivery.NewOrderGrpcMock()
	orderConn := new(grpc.ClientConn)

	grpcClient := grpcclient.NewGrpc(orderGrpc, orderConn)
	g.shippingCache = cache.NewShippingMock()

	g.shipperRestful = delivery.NewShipperMock()
	restfulClient := restfulclient.NewRestful(g.shipperRestful)

	g.shippingService = serviceimpl.NewShipping(restfulClient, grpcClient, g.shippingCache)
}

func (g *GetProvincesTestSuite) Test_Success() {
	shipperRes := g.CreateShipperRes()

	g.shippingCache.Mock.On("FindProvinces", mock.Anything).Return(nil)
	g.shipperRestful.Mock.On("GetProvinces", mock.Anything).Return(shipperRes, nil)

	res, err := g.shippingService.GetProvinces(context.Background())
	assert.NoError(g.T(), err)

	assert.Equal(g.T(), shipperRes, res)
}

func (g *GetProvincesTestSuite) Test_GetFromCache() {
	shipperRes := g.CreateShipperRes()

	g.shippingCache.Mock.On("FindProvinces", mock.Anything).Return(shipperRes)

	res, err := g.shippingService.GetProvinces(context.Background())
	assert.NoError(g.T(), err)

	assert.Equal(g.T(), shipperRes, res)
}

func (g *GetProvincesTestSuite) CreateShipperRes() *dto.ShipperRes[[]*entity.Province] {
	provinces := g.CreateProvinces()

	return &dto.ShipperRes[[]*entity.Province]{
		Data: provinces,
		Pagination: &entity.Pagination{
			CurrentPage:   1,
			TotalPages:    1,
			TotalElements: 34,
		},
	}
}

func (g *GetProvincesTestSuite) CreateProvinces() []*entity.Province {
	return []*entity.Province{
		{
			Id:   1,
			Name: "Bali",
			Lat:  -8.4095178,
			Lng:  115.188916,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   2,
			Name: "Bangka Belitung",
			Lat:  -2.7410513,
			Lng:  106.4405872,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   3,
			Name: "Banten",
			Lat:  -6.4058172,
			Lng:  106.0640179,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   4,
			Name: "Bengkulu",
			Lat:  -3.7928451,
			Lng:  102.2607641,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   5,
			Name: "DI Yogyakarta",
			Lat:  -7.7975915,
			Lng:  110.3707141,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   6,
			Name: "DKI Jakarta",
			Lat:  -6.1744651,
			Lng:  106.822745,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   7,
			Name: "Gorontalo",
			Lat:  0.5435442,
			Lng:  123.0567693,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   8,
			Name: "Jambi",
			Lat:  -1.6101229,
			Lng:  103.6131203,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   9,
			Name: "Jawa Barat",
			Lat:  -7.090911,
			Lng:  107.668887,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   10,
			Name: "Jawa Tengah",
			Lat:  -7.150975,
			Lng:  110.1402594,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   11,
			Name: "Jawa Timur",
			Lat:  -7.5360639,
			Lng:  112.2384017,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   12,
			Name: "Kalimantan Barat",
			Lat:  -0.2787808,
			Lng:  111.4752851,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   13,
			Name: "Kalimantan Selatan",
			Lat:  -3.0926415,
			Lng:  115.2837585,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   14,
			Name: "Kalimantan Tengah",
			Lat:  -1.6814878,
			Lng:  113.3823545,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   15,
			Name: "Kalimantan Timur",
			Lat:  0.5386586,
			Lng:  116.419389,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   16,
			Name: "Kalimantan Utara",
			Lat:  3.0730929,
			Lng:  116.0413889,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   17,
			Name: "Kepulauan Riau",
			Lat:  3.9456514,
			Lng:  108.1428669,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   18,
			Name: "Lampung",
			Lat:  -4.5585849,
			Lng:  105.4068079,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   19,
			Name: "Maluku Utara",
			Lat:  -3.2384616,
			Lng:  130.1452734,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   20,
			Name: "Maluku",
			Lat:  1.5709993,
			Lng:  127.8087693,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   21,
			Name: "Aceh",
			Lat:  4.695135,
			Lng:  96.7493993,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   22,
			Name: "Nusa Tenggara Barat",
			Lat:  -8.6529334,
			Lng:  117.3616476,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   23,
			Name: "Nusa Tenggara Timur",
			Lat:  -8.6573819,
			Lng:  121.0793705,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   24,
			Name: "Papua Barat",
			Lat:  -4.269928,
			Lng:  138.0803529,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   25,
			Name: "Papua",
			Lat:  -1.3361154,
			Lng:  133.1747162,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   26,
			Name: "Riau",
			Lat:  0.2933469,
			Lng:  101.7068294,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   27,
			Name: "Sulawesi Barat",
			Lat:  -2.8441371,
			Lng:  119.2320784,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   28,
			Name: "Sulawesi Selatan",
			Lat:  -3.6687994,
			Lng:  119.9740534,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   29,
			Name: "Sulawesi Tengah",
			Lat:  -1.4308543,
			Lng:  121.4456171,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   30,
			Name: "Sulawesi Tenggara",
			Lat:  -4.132111,
			Lng:  122.1746051,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   31,
			Name: "Sulawesi Utara",
			Lat:  1.3830312,
			Lng:  124.8410426,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   32,
			Name: "Sumatera Barat",
			Lat:  -0.7399397,
			Lng:  100.8000051,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   33,
			Name: "Sumatera Selatan",
			Lat:  -3.264013,
			Lng:  104.8882267,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
		{
			Id:   34,
			Name: "Sumatera Utara",
			Lat:  2.4853801,
			Lng:  99.5450974,
			Country: &entity.Country{
				Id:   228,
				Name: "INDONESIA",
				Code: "ID",
			},
		},
	}
}

func TestSetvice_GetProvinces(t *testing.T) {
	suite.Run(t, new(GetProvincesTestSuite))
}
