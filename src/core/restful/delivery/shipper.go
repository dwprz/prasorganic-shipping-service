package delivery

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dwprz/prasorganic-shipping-service/src/common/errors"
	"github.com/dwprz/prasorganic-shipping-service/src/common/helper"
	"github.com/dwprz/prasorganic-shipping-service/src/common/log"
	"github.com/dwprz/prasorganic-shipping-service/src/infrastructure/cbreaker"
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

func (s *ShipperImpl) ShippingOrder(ctx context.Context, data *entity.ShippingOrder) (shippingId string, err error) {
	res, err := cbreaker.Shipper.Execute(func() (any, error) {
		uri := config.Conf.Shipper.BaseUrl + "/v3/order"

		a := fiber.AcquireAgent()
		defer fiber.ReleaseAgent(a)

		a.JSON(data)

		req := a.Request()
		req.Header.SetContentType("application/json")
		req.Header.Set("X-API-KEY", config.Conf.Shipper.ApiKey)
		req.Header.SetMethod("POST")
		req.SetRequestURI(uri)

		if err := a.Parse(); err != nil {
			return "", err
		}

		code, body, _ := a.Bytes()
		if code != 201 {
			return "", &errors.Response{HttpCode: code, Message: string(body)}
		}

		res := new(struct {
			Data struct {
				ShippingId string `json:"order_id"`
			} `json:"data"`
		})

		err = json.Unmarshal(body, res)

		return res.Data.ShippingId, err
	})

	shippingId, ok := res.(string)
	if !ok {
		return "", fmt.Errorf("unexpected type %T expected string", res)
	}

	return shippingId, err
}

func (s *ShipperImpl) RequestPickup(ctx context.Context, shippingIds []string) error {
	_, err := cbreaker.Shipper.Execute(func() (any, error) {
		pickupReq := helper.FormatPickupReq(shippingIds)

		uri := config.Conf.Shipper.BaseUrl + "/v3/pickup"

		a := fiber.AcquireAgent()
		defer fiber.ReleaseAgent(a)

		a.JSON(pickupReq)

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
			return nil, &errors.Response{HttpCode: code, Message: string(body)}
		}

		log.Logger.Info(string(body)) // log success request pickup

		return nil, nil
	})

	return err
}

func (s *ShipperImpl) Pricing(ctx context.Context, data *dto.PricingReq) (*dto.ShipperRes[*entity.Pricing], error) {
	res, err := cbreaker.Shipper.Execute(func() (any, error) {
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
			return nil, &errors.Response{HttpCode: code, Message: string(body)}
		}

		res := new(dto.ShipperRes[*entity.Pricing])
		err := json.Unmarshal(body, res)

		return res, err
	})

	if err != nil {
		return nil, err
	}

	shipperRes, ok := res.(*dto.ShipperRes[*entity.Pricing])
	if !ok {
		return nil, fmt.Errorf("unexpected type %T expected *dto.ShipperRes[*entity.Pricing]", res)
	}

	return shipperRes, err
}

func (s *ShipperImpl) CreateLabel(ctx context.Context, data *dto.CreateLabelReq) (*dto.ShipperRes[*entity.Label], error) {
	res, err := cbreaker.Shipper.Execute(func() (any, error) {
		uri := config.Conf.Shipper.BaseUrl + "/v3/order/label"

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
		if code != 201 {
			return nil, &errors.Response{HttpCode: code, Message: string(body)}
		}

		res := new(dto.ShipperRes[*entity.Label])
		err := json.Unmarshal(body, res)

		return res, err
	})

	if err != nil {
		return nil, err
	}

	shipperRes, ok := res.(*dto.ShipperRes[*entity.Label])
	if !ok {
		return nil, fmt.Errorf("unexpected type %T expected *dto.ShipperRes[*entity.Label]", res)
	}

	return shipperRes, err
}

func (s *ShipperImpl) TrackingByShippingId(ctx context.Context, shippingId string) (*entity.Tracking, error) {
	res, err := cbreaker.Shipper.Execute(func() (any, error) {
		uri := fmt.Sprintf("%s/v3/order/%s", config.Conf.Shipper.BaseUrl, shippingId)
		a := fiber.AcquireAgent()
		defer fiber.ReleaseAgent(a)

		req := a.Request()
		req.Header.SetContentType("application/json")
		req.Header.Set("X-API-KEY", config.Conf.Shipper.ApiKey)
		req.Header.SetMethod("GET")
		req.SetRequestURI(uri)

		if err := a.Parse(); err != nil {
			return nil, err
		}

		code, body, _ := a.Bytes()
		if code != 200 {
			return nil, &errors.Response{HttpCode: code, Message: string(body)}
		}

		res := new(dto.ShipperRes[*entity.Tracking])
		err := json.Unmarshal(body, res)

		return res, err
	})

	if err != nil {
		return nil, err
	}

	shipperRes, ok := res.(*dto.ShipperRes[*entity.Tracking])
	if !ok {
		return nil, fmt.Errorf("unexpected type %T expected *dto.ShipperRes[*entity.Tracking]", res)
	}

	return shipperRes.Data, err
}

func (s *ShipperImpl) GetProvinces(ctx context.Context) (*dto.ShipperRes[[]*entity.Province], error) {
	res, err := cbreaker.Shipper.Execute(func() (any, error) {
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
			return nil, &errors.Response{HttpCode: code, Message: string(body)}
		}

		res := new(dto.ShipperRes[[]*entity.Province])
		err := json.Unmarshal(body, res)

		return res, err
	})

	if err != nil {
		return nil, err
	}

	shipperRes, ok := res.(*dto.ShipperRes[[]*entity.Province])
	if !ok {
		return nil, fmt.Errorf("unexpected type %T expected *dto.ShipperRes[[]*entity.Province]", res)
	}

	return shipperRes, err
}

func (s *ShipperImpl) GetCitiesByProvinceId(ctx context.Context, provinceId int) (*dto.ShipperRes[[]*entity.City], error) {
	res, err := cbreaker.Shipper.Execute(func() (any, error) {
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
			return nil, &errors.Response{HttpCode: code, Message: string(body)}
		}

		res := new(dto.ShipperRes[[]*entity.City])
		err := json.Unmarshal(body, res)

		return res, err
	})

	if err != nil {
		return nil, err
	}

	shipperRes, ok := res.(*dto.ShipperRes[[]*entity.City])
	if !ok {
		return nil, fmt.Errorf("unexpected type %T expected *dto.ShipperRes[[]*entity.City]", res)
	}

	return shipperRes, err
}

func (s *ShipperImpl) GetSuburbsByCityId(ctx context.Context, cityId int) (*dto.ShipperRes[[]*entity.Suburb], error) {
	res, err := cbreaker.Shipper.Execute(func() (any, error) {
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
			return nil, &errors.Response{HttpCode: code, Message: string(body)}
		}

		res := new(dto.ShipperRes[[]*entity.Suburb])
		err := json.Unmarshal(body, res)

		return res, err
	})

	if err != nil {
		return nil, err
	}

	shipperRes, ok := res.(*dto.ShipperRes[[]*entity.Suburb])
	if !ok {
		return nil, fmt.Errorf("unexpected type %T expected *dto.ShipperRes[[]*entity.Suburb]", res)
	}

	return shipperRes, err
}

func (s *ShipperImpl) GetAreasBySuburbId(ctx context.Context, suburbId int) (*dto.ShipperRes[[]*entity.Area], error) {
	res, err := cbreaker.Shipper.Execute(func() (any, error) {
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
			return nil, &errors.Response{HttpCode: code, Message: string(body)}
		}

		res := new(dto.ShipperRes[[]*entity.Area])
		err := json.Unmarshal(body, res)

		return res, err
	})

	if err != nil {
		return nil, err
	}

	shipperRes, ok := res.(*dto.ShipperRes[[]*entity.Area])
	if !ok {
		return nil, fmt.Errorf("unexpected type %T expected *dto.ShipperRes[[]*entity.Area]", res)
	}

	return shipperRes, err
}
