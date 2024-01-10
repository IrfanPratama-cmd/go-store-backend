package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetCategoryID(c *fiber.Ctx) error {
	db := services.DB
	id := c.Params("id")

	var category model.Category
	result := db.Model(&category).Where("id = ?", id).First(&category)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Category not found",
		})
	}

	return lib.OK(c, category)
}
