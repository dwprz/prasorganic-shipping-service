package delivery

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dwprz/prasorganic-shipping-service/src/infrastructure/config"
	"github.com/dwprz/prasorganic-shipping-service/src/interface/delivery"
	"github.com/dwprz/prasorganic-shipping-service/src/model/dto"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
	"github.com/gofiber/fiber/v2"
)

type ShipperImpl struct{}

func NewShipper() delivery.Shipper {
	return &ShipperImpl{}
}

func (s *ShipperImpl) Pricing(ctx context.Context, data *dto.PricingReq) (*dto.ShipperRes[*entity.Pricing], error) {
	uri := config.Conf.Shipper.BaseUrl + "/v3/pricing/domestic"

	a := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(a)
	
	a.JSON(data)

	req := a.Request()
	req.Header.SetContentType("application/json")
	req.Header.Set("X-API-KEY", config.Conf.Shipper.ApiKey)
	req.Header.SetMethod("POST")
	req.SetRequestURI(uri)

	if err := a.Parse(); err != nil {
		return nil, err
	}

	code, body, _ := a.Bytes()
	if code != 200 {
		return nil, fmt.Errorf(string(body))
	}

	res := new(dto.ShipperRes[*entity.Pricing])
	err := json.Unmarshal(body, res)

	return res, err
}

func (s *ShipperImpl) GetProvinces(ctx context.Context) (*dto.ShipperRes[[]*entity.Province], error) {
	uri := config.Conf.Shipper.BaseUrl + "/v3/location/country/228/provinces?limit=40"

	a := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(a)

	req := a.Request()
	req.Header.Set("X-API-KEY", config.Conf.Shipper.ApiKey)
	req.Header.SetMethod("GET")
	req.SetRequestURI(uri)

	if err := a.Parse(); err != nil {
		return nil, err
	}

	code, body, _ := a.Bytes()
	if code != 200 {
		return nil, fmt.Errorf(string(body))
	}

	res := new(dto.ShipperRes[[]*entity.Province])
	err := json.Unmarshal(body, res)

	return res, err
}

func (s *ShipperImpl) GetCitiesByProvinceId(ctx context.Context, provinceId int) (*dto.ShipperRes[[]*entity.City], error) {
	uri := fmt.Sprintf("%s/v3/location/province/%d/cities?limit=40", config.Conf.Shipper.BaseUrl, provinceId)

	a := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(a)

	req := a.Request()
	req.Header.Set("X-API-KEY", config.Conf.Shipper.ApiKey)
	req.Header.SetMethod("GET")
	req.SetRequestURI(uri)

	if err := a.Parse(); err != nil {
		return nil, err
	}

	code, body, _ := a.Bytes()
	if code != 200 {
		return nil, fmt.Errorf(string(body))
	}

	res := new(dto.ShipperRes[[]*entity.City])
	err := json.Unmarshal(body, res)

	return res, err
}

func (s *ShipperImpl) GetSuburbsByCityId(ctx context.Context, cityId int) (*dto.ShipperRes[[]*entity.Suburb], error) {
	uri := fmt.Sprintf("%s/v3/location/city/%d/suburbs?limit=51", config.Conf.Shipper.BaseUrl, cityId)

	a := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(a)

	req := a.Request()
	req.Header.Set("X-API-KEY", config.Conf.Shipper.ApiKey)
	req.Header.SetMethod("GET")
	req.SetRequestURI(uri)

	if err := a.Parse(); err != nil {
		return nil, err
	}

	code, body, _ := a.Bytes()
	if code != 200 {
		return nil, fmt.Errorf(string(body))
	}

	res := new(dto.ShipperRes[[]*entity.Suburb])
	err := json.Unmarshal(body, res)

	return res, err
}

func (s *ShipperImpl) GetAreasBySuburbId(ctx context.Context, suburbId int) (*dto.ShipperRes[[]*entity.Area], error) {
	uri := fmt.Sprintf("%s/v3/location/suburb/%d/areas?limit=35", config.Conf.Shipper.BaseUrl, suburbId)

	a := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(a)

	req := a.Request()
	req.Header.Set("X-API-KEY", config.Conf.Shipper.ApiKey)
	req.Header.SetMethod("GET")
	req.SetRequestURI(uri)

	if err := a.Parse(); err != nil {
		return nil, err
	}

	code, body, _ := a.Bytes()
	if code != 200 {
		return nil, fmt.Errorf(string(body))
	}

	res := new(dto.ShipperRes[[]*entity.Area])
	err := json.Unmarshal(body, res)

	return res, err
}
