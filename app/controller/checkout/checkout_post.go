package checkout

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2"
)

// PostCheckout godoc
// @Summary Create new Checkout
// @Description Create new Checkout
// @Accept  application/json
// @Produce application/json
// @Success 201 {object} model.Transaction "Checkout data"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /checkouts [post]
// @Tags Checkout
func PostCheckout(c *fiber.Ctx) error {
	db := services.DB

	userID := lib.GetXUserID(c)

	var contact model.Contact
	db.Model(&contact).Where("user_id", userID).First(&contact)

	var cart []model.Cart
	db.Model(&cart).Where("contact_id", contact.ID).Find(&cart)

	transactionID := lib.GenUUID()
	currentTime := time.Now()

	var totalAmount float64

	for _, c := range cart {
		var product model.Product
		db.Model(&product).Where("id", c.ProductID).First(&product)

		var checkout model.Checkout
		checkout.ContactID = contact.ID
		checkout.ProductID = c.ProductID
		checkout.TransactionID = transactionID
		checkout.Qty = c.Qty
		checkout.Amount = &product.Price
		checkout.TotalAmount = lib.Float64ptr(float64(*c.Qty) * product.Price)
		db.Create(&checkout)

		totalAmount += *checkout.TotalAmount
	}

	invoiceNo := lib.RandomNumber(6)

	var transaction model.Transaction
	transaction.ID = transactionID
	transaction.TotalAmount = &totalAmount
	transaction.TransactionDate = strfmt.DateTime(currentTime)
	transaction.TransactionStatus = lib.Strptr("pending")
	transaction.InvoiceNo = &invoiceNo
	transaction.ContactID = contact.ID
	db.Create(&transaction)

	db.Delete(&cart)

	return lib.Created(c, transaction)
}
