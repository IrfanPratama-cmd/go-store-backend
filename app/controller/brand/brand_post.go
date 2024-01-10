package brand

import (
	"api/app/model"
	"api/app/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PostBrand(c *fiber.Ctx) error {
	var brandAPI model.BrandAPI

	db := services.DB

	if err := c.BodyParser(&brandAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(brandAPI)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}
	brand := &model.Brand{BrandAPI: brandAPI}
	db.Model(&model.Brand{}).Create(brand)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    brand,
	})
}
