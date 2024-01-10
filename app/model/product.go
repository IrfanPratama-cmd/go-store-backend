package model

import "github.com/google/uuid"

type Product struct {
	Base
	ProductAPI
	Category     *Category     `json:"category,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
	Brand        *Brand        `json:"brand,omitempty" gorm:"foreignKey:BrandID;references:ID"`
	ProductAsset *ProductAsset `json:"product_asset,omitempty" gorm:"foreignKey:ProductID;references:ID"`
}

type ProductAPI struct {
	ProductName string     `json:"product_name,omitempty" validate:"required"`
	Description string     `json:"description,omitempty" validate:"required"`
	Quantity    int        `json:"quantity,omitempty"`
	Price       float64    `json:"price,omitempty"`
	CategoryID  *uuid.UUID `json:"category_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	BrandID     *uuid.UUID `json:"brand_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
}

type ProductRequest struct {
	ProductName  string         `json:"product_name,omitempty" validate:"required"`
	Description  string         `json:"description,omitempty" validate:"required"`
	Quantity     int            `json:"quantity,omitempty"`
	Price        float64        `json:"price,omitempty"`
	CategoryID   *uuid.UUID     `json:"category_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	BrandID      *uuid.UUID     `json:"brand_id,omitempty" gorm:"type:varchar(36);index" format:"uuid"`
	ProductAsset []ProductAsset `json:"product_asset,omitempty"`
}
