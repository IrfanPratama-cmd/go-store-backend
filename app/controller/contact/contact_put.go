package contact

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// PutContact godoc
// @Summary Update Contact by user id
// @Description Update Contact by user id
// @Param data body model.ContactRequest true "Contact data"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Contact "Contact data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /contacts/ [put]
// @Tags Contact
func PutContact(c *fiber.Ctx) error {
	api := new(model.ContactRequest)

	if err := c.BodyParser(&api); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	db := services.DB

	userID := lib.GetXUserID(c)
	var contact model.Contact
	db.Model(&contact).Where("user_id", userID).First(&contact)

	lib.Merge(api, &contact)

	db.Model(&contact).Updates(&contact)

	return lib.OK(c, contact)
}
