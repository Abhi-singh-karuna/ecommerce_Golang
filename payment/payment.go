package payment

import (
	"context"
	"os"
	"time"

	"github.com/Abhi-singh-karuna/Ronin/config"
	"github.com/Abhi-singh-karuna/Ronin/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func CashOnDelivery(c *fiber.Ctx) error {

	ProductCollection := config.MI.DB.Collection(os.Getenv("DATABASE_COLLECTION"))

	ProductCollection1 := config.MI.DB.Collection(os.Getenv("ORDER_DETAILS"))

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var order []models.OrderDeatails

	cursor, _ := ProductCollection.Find(ctx, bson.M{})

	Sum1 := 0

	for cursor.Next(ctx) {
		var product models.OrderDeatails

		cursor.Decode(&product)
		Sum1 = product.Price + Sum1

		ProductCollection1.InsertOne(ctx, models.OrderDeatails{
			Name:     product.Name,
			Quantity: product.Quantity,

			Price: product.Price,
		})

		order = append(order, product)

	}

	//remove the document for corresponding cart address
	//ProductCollection.Drop(ctx)

	//

	return c.JSON(fiber.Map{

		"total_price": Sum1,
		"message":     "thanks for shopping ",
		"details": fiber.Map{
			"data": order,
		},
	})

}
