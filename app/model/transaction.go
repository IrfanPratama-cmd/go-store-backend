package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type Transaction struct {
	Base
	TransactionAPI
	Contact *Contact `json:"contact,omitempty" gorm:"foreignKey:ContactID;references:ID"`
}

type TransactionAPI struct {
	ContactID         *uuid.UUID      `json:"contact_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	InvoiceNo         *string         `json:"invoice_no,omitempty"  gorm:"type:varchar(191)" example:"INV-000000000000072270623"`
	TransactionDate   strfmt.DateTime `json:"transaction_date,omitempty" format:"date-time" swaggertype:"string" gorm:"type:timestamptz;not null"`
	TransactionStatus *string         `json:"transaction_status,omitempty" gorm:"type:varchar(191);not null" example:"due" `
	TotalAmount       *float64        `json:"total_amount,omitempty" example:"127000"`
}

type TransactionRequest struct {
	OrderID         string `json:"order_id"`
	Amount          int    `json:"amount"`
	CreditCardToken string `json:"credit_card_token"`
}
