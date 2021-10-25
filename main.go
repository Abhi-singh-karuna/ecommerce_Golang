package main

import (
	c "github.com/KeshikaGupta20/Ronin_Cart/config"
	"github.com/KeshikaGupta20/Ronin_Cart/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	c.ConnectDB()
	routes.RegisterRoutes(app)

	app.Listen(":3001")
	//fmt.Println("Result of InsertOne")

}
