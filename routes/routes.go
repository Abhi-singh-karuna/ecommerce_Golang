package routes

import (
	"github.com/Abhi-singh-karuna/Ronin/authentication"
	"github.com/gofiber/fiber/v2"
)

func Route(app fiber.Router) {
	app.Post("/login", authentication.Login)
	app.Post("/signup", authentication.SignUp)
}
