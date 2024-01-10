package migrations

import "api/app/model"

// ModelMigrations models to automigrate
var ModelMigrations = []interface{}{
	&model.Asset{},
	&model.User{},
	&model.Contact{},
}
