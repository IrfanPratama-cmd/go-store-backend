package migrations

import "api/app/model"

var (
	category model.Category
	brand    model.Brand
)

// DataSeeds data to seeds
func DataSeeds() []interface{} {
	return []interface{}{
		category.Seed(),
		brand.Seed(),
	}
}
