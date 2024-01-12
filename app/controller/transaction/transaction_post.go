package transaction

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

// PostTransaction godoc
// @Summary Create Payment Transaction by Transaction id
// @Description Create Payment Transaction by Transaction id
// @Param id path string true "Transaction ID"
// @Accept  application/json
// @Produce application/json
// @Success 200 {object} model.PaymentLinkResponse "Payment Link Response"
// @Failure 400 {object} lib.Response
// @Failure 404 {object} lib.Response
// @Failure 409 {object} lib.Response
// @Failure 500 {object} lib.Response
// @Failure default {object} lib.Response
// @Security ApiKeyAuth
// @Router /transactions/{transaction_id} [post]
// @Tags Transaction
func PostTransaction(c *fiber.Ctx) error {
	// Inisialisasi library Xendit dan set API Key
	xendit.Opt.SecretKey = "xnd_development_CggMcaq5OxL4UuTU3iHobjHOBOP515vInD8bG0tc9oztmhxFzfLa5AUtCfUa5g"

	db := services.DB

	transaction_id := c.Params("transaction_id")

	var transaction model.Transaction
	db.Where("id = ?", transaction_id).First(&transaction)

	var contact model.Contact
	db.Where("id = ?", transaction.ContactID).First(&contact)

	var checkout []model.Checkout
	db.Model(&checkout).Preload("Product").Where("transaction_id = ?", transaction_id).Find(&checkout)

	var items []xendit.InvoiceItem

	for _, c := range checkout {
		product := c.Product

		item := xendit.InvoiceItem{
			Name:     product.ProductName,
			Quantity: *c.Qty,
			Price:    product.Price,
		}

		// Tambahkan item ke dalam slice items
		items = append(items, item)
	}

	createInvoiceData := invoice.CreateParams{
		ExternalID: *transaction.InvoiceNo,
		Amount:     *transaction.TotalAmount,
		PayerEmail: *contact.Email,
		Items:      items,
	}

	createdInvoice, err := invoice.Create(&createInvoiceData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create invoice"})
	}

	invoiceURL := "https://checkout-staging.xendit.co/latest/" + createdInvoice.ID

	var payment model.Payment
	payment.ExternalID = *transaction.InvoiceNo
	payment.Amount = *transaction.TotalAmount
	payment.PaymentStatus = "pending"
	payment.CheckoutLink = invoiceURL
	db.Create(&payment)

	return c.JSON(fiber.Map{"invoice_url": invoiceURL})

}

func XenditWebhookHandler(c *fiber.Ctx) error {
	// Mendapatkan data notifikasi dari Xendit
	webhookData := new(xendit.Invoice)

	if err := c.BodyParser(webhookData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	db := services.DB

	var payment model.Payment
	db.Where("external_id = ?", webhookData.ExternalID).First(&payment)
	payment.PaymentStatus = "settled"
	db.Where("external_id = ?", webhookData.ExternalID).Updates(&payment)

	var transaction model.Transaction
	db.Where("invoice_no = ?", webhookData.ExternalID).First(&transaction)
	transaction.TransactionStatus = lib.Strptr("paid")
	db.Where("invoice_no = ?", webhookData.ExternalID).Updates(&transaction)

	var checkout []model.Checkout
	db.Model(&checkout).Preload("Product").Where("transaction_id = ?", transaction.ID).Find(&checkout)

	for _, c := range checkout {
		product := c.Product

		product.Quantity -= *c.Qty
		db.Where("id = ?", product.ID).Updates(&product)
	}

	return c.JSON(fiber.Map{"message": "Webhook handled successfully"})
}
