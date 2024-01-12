package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
)

// GetTransactionID godoc
// @Summary Get a Transaction by id
// @Description Get a Transaction by id
// @Param id path string true "Transaction ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.Checkout "Transaction data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /transactions/{id} [get]
// @Tags Transaction
func GetTransactionID(c *fiber.Ctx) error {
	db := services.DB

	id := c.Params("id")

	userID := lib.GetXUserID(c)
	var contact model.Contact
	db.Model(&contact).Where("user_id", userID).First(&contact)

	var transaction model.Transaction
	result := db.Model(&transaction).
		Where(`transactions.contact_id = ?`, contact.ID).Where("id = ?", id).
		First(&transaction)

	var checkout []model.Checkout
	db.Model(&checkout).Preload("Product").Where("transaction_id = ?", id).Find(&checkout)

	if result.RowsAffected < 1 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Transaction ID Not Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"transaction": transaction,
		"checkout":    checkout,
	})

}
