package brand

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func PutBrand(c *fiber.Ctx) error {
	api := new(model.BrandAPI)

	if err := c.BodyParser(&api); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	id := c.Params("id")
	db := services.DB

	var data model.Brand
	result := db.Model(&data).Where("id = ?", &id).Take(&data)

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Brand Not Found",
		})
	}

	lib.Merge(api, &data)

	db.Model(&data).Updates(&data)

	return lib.OK(c, data)
}
