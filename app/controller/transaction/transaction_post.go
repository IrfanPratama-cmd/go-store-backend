package transaction

import (
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

func PostTransaction(c *fiber.Ctx) error {
	// var request model.PaymentAPI

	db := services.DB

	transaction_id := c.Params("transaction_id")

	var transaction model.Transaction
	db.Model(&transaction).Where("id = ?", transaction_id).First(&transaction)

	var contact model.Contact
	db.Model(&contact).Where("id = ?", transaction.ContactID).First(&contact)

	// Inisialisasi library Xendit dan set API Key
	xendit.Opt.SecretKey = "xnd_development_CggMcaq5OxL4UuTU3iHobjHOBOP515vInD8bG0tc9oztmhxFzfLa5AUtCfUa5g"

	createInvoiceData := invoice.CreateParams{
		ExternalID: transaction_id,
		Amount:     *transaction.TotalAmount,
		PayerEmail: *contact.Email,
	}

	createdInvoice, err := invoice.Create(&createInvoiceData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create invoice"})
	}

	invoiceURL := "https://checkout-staging.xendit.co/latest/" + createdInvoice.ID

	var payment model.Payment
	payment.ExternalID = transaction_id
	payment.Amount = *transaction.TotalAmount
	payment.PaymentStatus = "Pending"
	payment.CheckoutLink = invoiceURL
	db.Create(&payment)

	return c.JSON(fiber.Map{"invoice_url": invoiceURL})

}
