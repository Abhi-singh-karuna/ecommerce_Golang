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

	//initiate the fiber framework
	app := fiber.New()

	//use logger for detailed view
	app.Use(logger.New())

	//for intialiseing thr env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//configure your database
	config.ConnectDB()

	//call the setupRoutes route
	setRoutes(app)

	//listen by this port
	err = app.Listen(":9000")
	if err != nil {
		log.Fatal(err)
	}

}



func setRoutes(app *fiber.App) {
	//if you want to group up your Url
	api := app.Group("/")

	//call the route package
	routes.BookRutes(api.Group("/product"))

}
