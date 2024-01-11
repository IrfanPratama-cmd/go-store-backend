package migrations

import "api/app/model"

// ModelMigrations models to automigrate
var ModelMigrations = []interface{}{
	&model.Asset{},
	&model.User{},
	&model.Contact{},
	&model.Category{},
	&model.Brand{},
	&model.Product{},
	&model.ProductAsset{},
	&model.Checkout{},
	&model.Cart{},
	&model.Transaction{},
	&model.Payment{},
}
