package handler

import (
	"strconv"

	"github.com/dwprz/prasorganic-shipping-service/src/interface/service"
	"github.com/dwprz/prasorganic-shipping-service/src/model/dto"
	"github.com/dwprz/prasorganic-shipping-service/src/model/entity"
	"github.com/gofiber/fiber/v2"
)

type Shipping struct {
	shippingService service.Shipping
}

func NewShipping(ss service.Shipping) *Shipping {
	return &Shipping{
		shippingService: ss,
	}
}

func (s *Shipping) ManualShipping(c *fiber.Ctx) error {
	req := new(entity.ShippingOrder)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	err := s.shippingService.ShippingOrder(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"data": "order shipped successfully"})
}

func (s *Shipping) CreateLabel(c *fiber.Ctx) error {
	req := new(dto.CreateLabelReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := s.shippingService.CreateLabel(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"data": res})
}

func (s *Shipping) RequestPickup(c *fiber.Ctx) error {
	var req map[string][]string

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err := s.shippingService.RequestPickup(c.Context(), req["shipping_ids"])
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"data": "pickup successfully requested"})
}

func (s *Shipping) Tracking(c *fiber.Ctx) error {
	shippingId := c.Params("shippingId")

	res, err := s.shippingService.TrackingByShippingId(c.Context(), shippingId)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res})
}

func (s *Shipping) GetProvinces(c *fiber.Ctx) error {
	res, err := s.shippingService.GetProvinces(c.Context())
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res.Data, "pagination": res.Pagination})
}

func (s *Shipping) GetCities(c *fiber.Ctx) error {
	provinceId, err := strconv.Atoi(c.Query("provinceId"))
	if err != nil {
		return err
	}

	res, err := s.shippingService.GetCitiesByProvinceId(c.Context(), provinceId)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res.Data, "pagination": res.Pagination})
}

func (s *Shipping) GetSuburbs(c *fiber.Ctx) error {
	cityId, err := strconv.Atoi(c.Query("cityId"))
	if err != nil {
		return err
	}

	res, err := s.shippingService.GetSuburbsByCityId(c.Context(), cityId)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res.Data, "pagination": res.Pagination})
}

func (s *Shipping) GetAreas(c *fiber.Ctx) error {
	cityId, err := strconv.Atoi(c.Query("suburbId"))
	if err != nil {
		return err
	}

	res, err := s.shippingService.GetAreasBySuburbId(c.Context(), cityId)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res.Data, "pagination": res.Pagination})
}

func (s *Shipping) Pricing(c *fiber.Ctx) error {
	req := new(dto.PricingReq)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := s.shippingService.Pricing(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res.Data, "pagination": res.Pagination})
}
