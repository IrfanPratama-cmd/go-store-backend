package category

import (
	"api/app/model"
	"api/app/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func PostCategory(c *fiber.Ctx) error {
	var categoryAPI model.CategoryAPI

	db := services.DB

	if err := c.BodyParser(&categoryAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(categoryAPI)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}
	category := &model.Category{CategoryAPI: categoryAPI}
	db.Model(&model.Category{}).Create(category)

	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    category,
	})
}
