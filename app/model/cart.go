package model

import "github.com/google/uuid"

type Cart struct {
	Base
	CartAPI
	Contact *Contact `json:"contact,omitempty" gorm:"foreignKey:ContactID;references:ID"`
	Product *Product `json:"product,omitempty" gorm:"foreignKey:ProductID;references:ID"`
}

type CartAPI struct {
	ContactID *uuid.UUID `json:"contact_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	ProductID *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	Qty       *int       `json:"qty,omitempty"`
}

type CartRequest struct {
	ProductID *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	Qty       *int       `json:"qty,omitempty"`
}

type CartUpdate struct {
	Qty *int `json:"qty,omitempty"`
}
