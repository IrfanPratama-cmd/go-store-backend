package brand

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func DeleteBrand(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	var brand model.Brand
	result := db.Model(&brand).Where("id = ?", id).First(&brand)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Brand not found",
		})
	}

	db.Delete(&brand)

	return lib.OK(c)
}
