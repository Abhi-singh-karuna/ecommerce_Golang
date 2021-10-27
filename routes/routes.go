package routes

import (
	c "github.com/KeshikaGupta20/Ronin_Cart/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app fiber.Router) {

	app.Post("/insertpro", c.CreateProduct)
	app.Delete("/deletepro/:id", c.DeleteProduct)
	app.Get("/getpro", c.GetProduct)
	app.Put("/updatepro/:id",c.UpdateProduct)
	app.Get("/getproid/:id", c.GetProductbyid)
}
