package cart

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
)

// GetCart godoc
// @Summary List of Cart
// @Description List of Cart
// @Param page query int false "Page number start from zero"
// @Param size query int false "Size per page, default `0`"
// @Param sort query string false "Sort by field, adding dash (`-`) at the beginning means descending and vice versa"
// @Param fields query string false "Select specific fields with comma separated"
// @Param filters query string false "custom filters, see [more details](https://github.com/morkid/paginate#filter-format)"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} lib.Page{items=[]model.Cart} "List of Cart"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /carts [get]
// @Tags Cart
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
