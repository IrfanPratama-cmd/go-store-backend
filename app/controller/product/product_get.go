package product

import (
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

func GetProduct(c *fiber.Ctx) error {
	db := services.DB
	pg := paginate.New()

	var product model.Product

	mod := db.Model(&product).
		Select(`products.id, products.created_at, products.updated_at, 
				products.product_name, products.description, products.price, products.quantity,
				products.thumbnail,
				c.id "Category__id",
				c.category_name "Category__category_name",
				c.category_code "Category__category_code",
				b.id "Brand__id",
				b.brand_name "Brand__brand_name",
				b.brand_code "Brand__brand_code",
				pa.id "ProductAsset__id",
				pa.file_name "ProductAsset__file_name",
				pa.file_path "ProductAsset__file_path",
				pa.is_primary "ProductAsset__is_primary"
				`).
		Joins(`LEFT JOIN brands b on b.id = products.brand_id`).
		Joins(`LEFT JOIN categories c on c.id = products.category_id`).
		Joins(`LEFT JOIN product_assets pa on pa.product_id = products.id`).
		Where(`pa.is_primary = ?`, true)

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Product{})
	return c.Status(200).JSON(page)
}
