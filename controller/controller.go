package controller

import (
	"os"

	"github.com/Abhi-singh-karuna/Ronin/config"
	"github.com/Abhi-singh-karuna/Ronin/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateData(c *fiber.Ctx) error {

	ConnectDb := config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION"))

	data := new(models.Product)

	c.BodyParser(&data)

	//data.ID = nil

	result, _ := ConnectDb.InsertOne(c.Context(), data)

	////////////////////////////////////////////////

	content := &models.Product{}

	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	ConnectDb.FindOne(c.Context(), query).Decode(content)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"meassage": "successfull",
		"data": fiber.Map{
			"data": content,
		},
	})

}
