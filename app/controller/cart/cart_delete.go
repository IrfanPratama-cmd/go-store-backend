package cart

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteCart(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	var cart model.Cart
	result := db.Model(&cart).Where("id = ?", id).First(&cart)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Cart not found",
		})
	}

	db.Delete(&cart)

	return lib.OK(c)
}
