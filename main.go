package main

import (
	"log"

	"github.com/Abhi-singh-karuna/Ronin/config"
	"github.com/Abhi-singh-karuna/Ronin/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())

	//for intialiseing thr env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	setRoutes(app)

	err = app.Listen(":9000")
	if err != nil {
		log.Fatal(err)
	}

}

func setRoutes(app *fiber.App) {

	api := app.Group("/")

	routes.BookRutes(api.Group("/product"))

}
