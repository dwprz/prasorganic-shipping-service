package router

import (
	"github.com/dwprz/prasorganic-shipping-service/src/core/restful/handler"
	"github.com/dwprz/prasorganic-shipping-service/src/core/restful/middleware"
	"github.com/gofiber/fiber/v2"
)

func Create(app *fiber.App, h *handler.Shipping, m *middleware.Middleware) {
	// admin & super admin
	app.Add("POST", "/api/shippings/orders", m.VerifyJwt, m.VerifyAdmin, h.ManualShipping)

	// all
	app.Add("POST", "/api/shippings/pricings", m.VerifyJwt, h.Pricing)
	app.Add("GET", "/api/shippings/:shippingId/trackings", m.VerifyJwt, h.Tracking)
	app.Add("GET", "/api/shippings/provinces", m.VerifyJwt, h.GetProvinces)
	app.Add("GET", "/api/shippings/cities", m.VerifyJwt, h.GetCities)
	app.Add("GET", "/api/shippings/suburbs", m.VerifyJwt, h.GetSuburbs)
	app.Add("GET", "/api/shippings/areas", m.VerifyJwt, h.GetAreas)
}