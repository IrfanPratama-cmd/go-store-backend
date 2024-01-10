package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

func GetCategory(c *fiber.Ctx) error {
	db := services.DB
	pg := paginate.New()

	mod := db.Model(&model.Category{})

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Category{})

	return lib.OK(c, page)
}
