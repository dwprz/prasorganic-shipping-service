package handler

import (
	"github.com/dwprz/prasorganic-shipping-service/src/interface/service"
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

func (s *Shipping) GetProvinces(c *fiber.Ctx) error {
	res, err := s.shippingService.GetProvinces(c.Context())
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"data": res})
}
