package product

import (
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

func GetProductID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := services.DB

	var product model.Product
	result := db.Model(&product).
		Select(`products.id, products.created_at, products.updated_at, 
				products.product_name, products.description, products.price, products.quantity,
				products.thumbnail,
				c.id "Category__id",
				c.category_name "Category__category_name",
				c.category_code "Category__category_code",
				b.id "Brand__id",
				b.brand_name "Brand__brand_name",
				b.brand_code "Brand__brand_code"
				`).
		Joins(`LEFT JOIN brands b on b.id = products.brand_id`).
		Joins(`LEFT JOIN categories c on c.id = products.category_id`).
		Where(`products.id = ?`, id).
		First(&product)

	if result.RowsAffected < 1 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product ID Not Found",
		})
	}

	var asset []model.ProductAsset
	db.Where(`product_id = ?`, product.ID).Find(&asset)
	return c.Status(200).JSON(fiber.Map{
		"product":       product,
		"product_asset": asset,
	})
}
