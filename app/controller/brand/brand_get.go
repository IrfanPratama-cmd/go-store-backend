package brand

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

func GetBrand(c *fiber.Ctx) error {
	db := services.DB
	pg := paginate.New()

	mod := db.Model(&model.Brand{})

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Brand{})

	return lib.OK(c, page)
}
