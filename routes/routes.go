package routes

import (
	"github.com/Abhi-singh-karuna/Ronin/Order"
	"github.com/Abhi-singh-karuna/Ronin/address"
	"github.com/Abhi-singh-karuna/Ronin/payment"

	"github.com/gofiber/fiber/v2"
)

func BookRutes(route fiber.Router) {
	route.Get("", Order.GetAllProductList)
	route.Get("/cart", Order.GetPriceAndProductName)
	route.Get("/cod", payment.CashOnDelivery1)

	route.Post("/address", address.AddAddress)
	route.Get("/address/:id", address.GetOneAddress)
	route.Get("/address", address.GetAllAddress)
	route.Put("/address/:id", address.UpdateAddress)
	route.Delete("/address/:id", address.DeleteAddress)

}
