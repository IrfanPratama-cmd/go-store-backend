package category

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

// GetCategory godoc
// @Summary List of Category
// @Description List of Category
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Page{items=[]model.Category} "List of Category"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /categories [get]
// @Tags Category
func GetCategory(c *fiber.Ctx) error {
	db := services.DB
	pg := paginate.New()

	mod := db.Model(&model.Category{})

	page := pg.With(mod).Request(c.Request()).Response(&[]model.Category{})

	return lib.OK(c, page)
}
