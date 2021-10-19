package Order

import (
	"os"

	"github.com/Abhi-singh-karuna/Ronin/config"
	"github.com/Abhi-singh-karuna/Ronin/models"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber/v2"
)

func GetAllProductList(c *fiber.Ctx) error {

	ProductCollection := config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION"))

	query := bson.D{{}}

	cursor, err := ProductCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	var Products []models.Product = make([]models.Product, 0)

	err = cursor.All(c.Context(), &Products)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"Products": Products,
		},
	})

}

func GetPriceAndProductName(c *fiber.Ctx) error {

	ProductCollection := config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION"))

	query := bson.D{{}}

	cursor, err := ProductCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}
	//	config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION")).find().forEach(function(d){ db.getSiblingDB('os.Getenv("ORDER_DETAILS")')['os.Getenv("DATABASE_COLLECTION")'].insert(d); });

	var Products []models.ProductCart = make([]models.ProductCart, 0)

	err = cursor.All(c.Context(), &Products)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"products": Products,
		},
	})

}
