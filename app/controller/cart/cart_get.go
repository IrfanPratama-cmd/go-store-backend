package cart

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

func GetCart(c *fiber.Ctx) error {
	db := services.DB
	pg := paginate.New()

	userID := lib.GetXUserID(c)
	var contact model.Contact
	db.Model(&contact).Where("user_id", userID).First(&contact)

	mod := db.Model(&model.Cart{}).Preload("Product").Preload("Product.ProductAsset").Where("contact_id = ?", contact.ID)

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Cart{})

	return lib.OK(c, page)

}
