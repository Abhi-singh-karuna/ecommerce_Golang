package routes

import (
	"github.com/Abhi-singh-karuna/Ronin/Order"
	"github.com/Abhi-singh-karuna/Ronin/address"
	"github.com/Abhi-singh-karuna/Ronin/payment"

	"github.com/gofiber/fiber/v2"
)

func BookRutes(route fiber.Router) {
	//route.Get("", Order.GetAllProductList) // fetch all the products stored in database
	route.Get("/cart", Order.GetPriceAndProductName) // fetch the data those stoed in cart collection
	route.Get("/cod", payment.CashOnDelivery) // for pacing the order in cash on delivery

	route.Post("/address", address.AddAddress) // take shipping address from users
	route.Get("/address/:id", address.GetOneAddress) // show the selected one address
	route.Get("/address", address.GetAllAddress)  // show all the address saved in particular user database
	route.Put("/address/:id", address.UpdateAddress) //update the address
	route.Delete("/address/:id", address.DeleteAddress) // delete the address
	route.Get("/filter", address.WorkOnData)


}
