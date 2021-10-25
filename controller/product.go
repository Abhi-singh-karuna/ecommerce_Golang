package controller

import (
	"github.com/KeshikaGupta20/Ronin_Cart/config"
	"github.com/KeshikaGupta20/Ronin_Cart/models"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber/v2"
)

// controller to get all the book from data slice

func GetProductList(c *fiber.Ctx) error {

	// Collection gets a handle for a collection with the given name configured with the given CollectionOptions
	ProductCollection := config.MI.DB.Collection("product")

	// Query to filter or iterate over the all the element
	query := bson.D{{}}

	// Context returns *fasthttp.RequestCtx that carries a deadline
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
		"Message": "Product inserted sucessfully",
	})

}
