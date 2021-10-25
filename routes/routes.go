package routes

import (
	c "github.com/KeshikaGupta20/Ronin_Cart/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app fiber.Router) {

	app.Post("/insertproduct", c.GetProductList)
	app.Get("/getproduct",  )
}
