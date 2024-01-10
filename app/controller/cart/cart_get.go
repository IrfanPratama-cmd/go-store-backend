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

	mod := db.Model(&model.Cart{}).Preload("Product").Where("user_id = ?", userID)

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Cart{})

	return lib.OK(c, page)

}
