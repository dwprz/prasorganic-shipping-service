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

func (s *ShipperImpl) GetProvinces(ctx context.Context) (*dto.ShipperRes[[]*entity.Province], error) {
	uri := config.Conf.Shipper.BaseUrl + "/v3/location/country/228/provinces?limit=40"

	a := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(a)

	req := a.Request()
	req.Header.Set("X-API-KEY", config.Conf.Shipper.ApiKey)
	req.SetRequestURI(uri)

	if err := a.Parse(); err != nil {
		return nil, err
	}

	_, body, errs := a.Bytes()
	if len(errs) > 0 {
		return nil, errs[0]
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
	req.SetRequestURI(uri)

	if err := a.Parse(); err != nil {
		return nil, err
	}

	_, body, errs := a.Bytes()
	if len(errs) > 0 {
		return nil, errs[0]
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
	req.SetRequestURI(uri)

	if err := a.Parse(); err != nil {
		return nil, err
	}

	_, body, errs := a.Bytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	res := new(dto.ShipperRes[[]*entity.Suburb])
	err := json.Unmarshal(body, res)

	return res, err
}
