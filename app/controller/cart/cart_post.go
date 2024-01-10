package cart

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func PostCart(c *fiber.Ctx) error {
	var cartAPI model.CartAPI

	db := services.DB

	if err := c.BodyParser(&cartAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	userID := lib.GetXUserID(c)

	var contact model.Contact
	db.Model(&contact).Where("user_id", userID).First(&contact)

	var cart model.Cart
	cart.ProductID = cartAPI.ProductID
	cart.Qty = cartAPI.Qty
	cart.ContactID = contact.ID
	db.Create(&cart)

	return lib.Created(c, cart)
}
