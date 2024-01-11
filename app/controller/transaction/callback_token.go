package transaction

import (
	"api/app/lib"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Webhook(c *fiber.Ctx) error {
	// Proses notifikasi pembayaran di sini
	var webhookData map[string]interface{}
	if err := c.BodyParser(&webhookData); err != nil {
		fmt.Println("Error parsing webhook data:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	fmt.Println("Received Payment Webhook:", webhookData)
	// Proses data webhook sesuai kebutuhan

	// return c.SendStatus(fiber.StatusOK)
	return lib.OK(c, webhookData)

}
