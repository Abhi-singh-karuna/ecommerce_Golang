package routes

import (
	"github.com/Abhi-singh-karuna/Ronin/controller"
	"github.com/gofiber/fiber/v2"
)

func CallRoutes(app fiber.Router) {

	app.Post("/", controller.CreateData)

}
