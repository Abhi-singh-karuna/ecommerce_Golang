package address

import (
	"context"
	"math"
	"strconv"
	"time"

	"fmt"
	"os"

	"github.com/Abhi-singh-karuna/Ronin/config"
	"github.com/Abhi-singh-karuna/Ronin/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddAddress(c *fiber.Ctx) error {

	AddressCollection := config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION_ADDRESS"))

	data := new(models.Add)

	err := c.BodyParser(&data)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})

	}
	data.ID = nil

	result, err := AddressCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert todo",
			"error":   err,
		})
	}

	Address := &models.Add{}

	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	AddressCollection.FindOne(c.Context(), query).Decode(Address)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Address": fiber.Map{
			"data": Address,
		},
	})
}

func GetAllAddress(c *fiber.Ctx) error {

	AddressCollection := config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION_ADDRESS"))

	query := bson.D{{}}

	cursor, err := AddressCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	var Address []models.Add = make([]models.Add, 0)

	err = cursor.All(c.Context(), &Address)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Address": fiber.Map{
			"book": Address,
		},
	})

}

func GetOneAddress(c *fiber.Ctx) error {

	AddressCollection := config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION_ADDRESS"))

	Addresid := c.Params("id")

	id, _ := primitive.ObjectIDFromHex(Addresid)

	Address := &models.Add{}

	query := bson.D{{Key: "_id", Value: id}}

	err := AddressCollection.FindOne(c.Context(), query).Decode(Address)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Todo Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"Address": Address,
		},
	})

}

func UpdateAddress(c *fiber.Ctx) error {

	AddressCollection := config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION_ADDRESS"))

	paramID := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}
	data := new(models.Add)
	err = c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	query := bson.D{{Key: "_id", Value: id}}

	var dataToUpdate bson.D

	if data.Full_Name != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "full_name", Value: data.Full_Name})
	}

	if data.Address_a != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "address_a", Value: data.Address_a})
	}

	if data.Address_b != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "address_b", Value: data.Address_b})
	}
	if data.City != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "city", Value: data.City})
	}
	if data.State != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "state", Value: data.State})
	}
	if data.PinCode != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "pincode", Value: data.PinCode})
	}

	update := bson.D{
		{Key: "$set", Value: dataToUpdate},
	}

	// update
	err = AddressCollection.FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Todo Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update todo",
			"error":   err,
		})
	}

	// get updated data
	Address := &models.Add{}

	AddressCollection.FindOne(c.Context(), query).Decode(Address)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Address": fiber.Map{
			"todo": Address,
		},
	})

}

func DeleteAddress(c *fiber.Ctx) error {

	AddressCollection := config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION_ADDRESS"))

	Addressid := c.Params("id")

	id, err := primitive.ObjectIDFromHex(Addressid)

	fmt.Println(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "cannot parse id",
			"error":   err,
		})
	}

	query := bson.D{{Key: "_id", Value: id}}

	err = AddressCollection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Address Not found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete todo",
			"error":   err,
		})
	}

	return c.SendString(" deletion succressfull ")

}

func WorkOnData(c *fiber.Ctx) error {
	collection := config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION_ADDRESS"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var products []models.Add

	filter := bson.M{}

	findOptions := options.Find()

	if s := c.Query("s"); s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"city": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
				{
					"state": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))

	var perpage int64 = 4

	total, _ := collection.CountDocuments(ctx, filter)

	findOptions.SetSkip((int64(page) - 1) * perpage)
	findOptions.SetLimit(int64(perpage))

	cursor, _ := collection.Find(ctx, filter, findOptions)

	for cursor.Next(ctx) {
		var product models.Add

		cursor.Decode(&product)

		products = append(products, product)
	}
	return c.JSON(fiber.Map{
		"data":          products,
		"total_Address": total,
		"page":          page,
		"last_page":     math.Ceil(float64(total / perpage)),
	})

}
