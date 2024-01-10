package model

import "github.com/google/uuid"

type Checkout struct {
	Base
	CheckoutAPI
	Contact    *Contact    `json:"contact,omitempty" gorm:"foreignKey:ContactID;references:ID"`
	Transacion *Transacion `json:"transaction,omitempty" gorm:"foreignKey:TransactionID;references:ID"`
}

type CheckoutAPI struct {
	ContactID     *uuid.UUID `json:"contact_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	ProductID     *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	TransactionID *uuid.UUID `json:"transaction_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	Qty           *int       `json:"qty,omitempty"`
	Amount        *float64   `json:"amount,omitempty"`
}
