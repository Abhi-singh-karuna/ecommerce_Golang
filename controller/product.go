package controller

import (
	"fmt"

	"github.com/KeshikaGupta20/Ronin_Cart/config"
	"github.com/KeshikaGupta20/Ronin_Cart/models"

	//"golang.org/x/text/message"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// controller to get all the book from data slice

func CreateProduct(c *fiber.Ctx) error {

	// Collection gets a handle for a collection with the given name configured with the given CollectionOptions
	ProductCollection := config.MI.DB.Collection("product")

	Pro := new(models.Product)

	c.BodyParser(&Pro)

	ProductCollection.InsertOne(c.Context(), Pro)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"Message": "Product inserted sucessfully",
	})

}
func DeleteProduct(c *fiber.Ctx) error {

	ProductCollection := config.MI.DB.Collection("product")

	ID := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(ID)

	if err != nil {

		return err
	}
	data := &models.Product{}

	filter := bson.D{{Key: "_id", Value: objectId}}

	ProductCollection.FindOne(c.Context(), filter).Decode(data)

	_, err = ProductCollection.DeleteOne(c.Context(), data)

	if err != nil {

		fmt.Println("Database Connected")

		return err
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"Message": "Product deleted sucessfully",
	})
}

func GetProduct(c *fiber.Ctx) error {

	ProductCollection := config.MI.DB.Collection("product")

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

func GetProductbyid(c *fiber.Ctx) error {

	ProductCollection := config.MI.DB.Collection("product")

	ID := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}
	data := &models.Product{}

	filter := bson.D{{Key: "_id", Value: objectId}}

	ProductCollection.FindOne(c.Context(), filter).Decode(data)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"data": fiber.Map{
			"Products": data,
		},
	})

}

func UpdateProduct(c *fiber.Ctx) error {

	ProductCollection := config.MI.DB.Collection("product")

	ID := c.Params("id")

	objectId, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}
	data := new(models.Product)
	c.BodyParser(&data)

	filter := bson.D{{Key: "_id", Value: objectId}}

	var message bson.D

	message = append(message, bson.E{Key: "name", Value: data.Name},
		bson.E{Key: "price", Value: data.Price}, bson.E{Key: "Quantity", Value: data.Quantity},
		bson.E{Key: "Description", Value: data.Description}, bson.E{Key: "image", Value: data.Image})

	update := bson.D{{"$set", message}}
	/* update := bson.D{
	{"$inc", bson.D{
		{"name", "sony smart television"},
	}}} */
	/* update := bson.D{{
		Key:   "name",
		Value: data.Name,
	}} */
	//ctx, bson.M{"email": user.Email}
	ProductCollection.UpdateOne(c.Context(), filter, update)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"success": true,
		"data": fiber.Map{
			"name":        data.Name,
			"price":       data.Price,
			"Quantity":    data.Quantity,
			"Description": data.Description,
			"image":       data.Image,
		},
	})
}
