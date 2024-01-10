package model

import (
	"github.com/google/uuid"
)

// Contact Contact
type Contact struct {
	Base
	DataOwner
	ContactAPI
}

// ContactAPI Contact API
type ContactAPI struct {
	ContactName     *string    `json:"contact_name,omitempty" example:"Walk-in-customers" gorm:"type:varchar(191);not null" validate:"required"`                                               // Contact Name                                                                                // Type
	Mobile          *string    `json:"mobile,omitempty" example:"08123456789" gorm:"type:varchar(191)"`                                                                                        // Mobile
	AlternateNumber *string    `json:"alternate_number,omitempty" example:"08123456789" gorm:"type:varchar(191)"`                                                                              // Alternate Number
	Email           *string    `json:"email,omitempty" example:"walk-in-customer@mail.com" gorm:"type:varchar(191)"`                                                                           // Email
	Website         *string    `json:"website,omitempty" example:"www.walk-in-customer.com" gorm:"type:varchar(191)"`                                                                          // Website
	ProvinceID      *uuid.UUID `json:"province_id,omitempty" swaggertype:"string" format:"uuid"`                                                                                               // Province ID
	CityID          *uuid.UUID `json:"city_id,omitempty" swaggertype:"string" format:"uuid"`                                                                                                   // City ID
	SubdistrictID   *uuid.UUID `json:"subdistrict_id,omitempty" swaggertype:"string" format:"uuid"`                                                                                            // Subdistrict ID
	Address         *string    `json:"address,omitempty" example:"Jl. Aria Putra No.88, RT.09/RW.01, Sawah Baru, Kec. Ciputat, Kota Tangerang Selatan, Banten 15414" gorm:"type:varchar(255)"` // Address
	ZipCode         *string    `json:"zip_code,omitempty" example:"15414" gorm:"type:varchar(255)"`                                                                                            // Zip Code
}

// InitialSetup Data
// func (b *Contact) InitialSetup(*uuid.UUID, tx *gorm.DB) error {
// 	types := "customer"
// 	contactName := "Walk-in-customers"
// 	b.Type = &types
// 	b.ContactName = &contactName
// 	if err := tx.Create(&b).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
