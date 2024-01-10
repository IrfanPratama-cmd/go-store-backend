package product

import (
	"api/app/model"
	"api/app/services"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func PutProduct(c *fiber.Ctx) error {
	var productAPI model.ProductAPI
	if err := c.BodyParser(&productAPI); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	id := c.Params("id")
	db := services.DB
	validate := validator.New()
	errValidate := validate.Struct(productAPI)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// thumbnail := c.Locals("thumbnail").(string)
	// asset := fmt.Sprintf("%v", thumbnail)

	var product model.Product

	if rowsAffected := db.First(&product, `id = ?`, id).RowsAffected; rowsAffected < 1 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product ID Not Found",
		})
	} else {
		product.ProductName = productAPI.ProductName
		product.Description = productAPI.Description
		product.Price = productAPI.Price
		product.Quantity = productAPI.Quantity
		product.CategoryID = productAPI.CategoryID
		// product.Thumbnail = thumbnail
		db.Model(&product).Where(`id = ?`, id).Updates(&product)
	}

	message := fmt.Sprintf(`Product with id %s has been updated`, id)

	return c.Status(200).JSON(fiber.Map{
		"message": message,
		"data":    product,
	})

}
