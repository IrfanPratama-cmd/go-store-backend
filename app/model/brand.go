package model

import "strings"

type Brand struct {
	Base
	BrandAPI
}

type BrandAPI struct {
	BrandCode *string `json:"brand_code,omitempty" example:"XM" validate:"required" gorm:"unique"`
	BrandName *string `json:"brand_name,omitempty" example:"Iphone" validate:"required" gorm:"unique"`
}

func (s Brand) Seed() *[]Brand {
	data := []Brand{}
	items := []string{
		"AS|Asus",
		"SM|Samsung",
		"XM|Xiaomi",
		"IP|Iphone",
		"AC|Acer",
		"OP|Oppo",
		"VO|Vivo",
	}

	for i := range items {
		var content string = items[i]
		c := strings.Split(content, "|")
		brandCode := c[0]
		brandName := c[1]

		data = append(data, Brand{
			BrandAPI: BrandAPI{
				BrandCode: &brandCode,
				BrandName: &brandName,
			},
		})
	}
	return &data
}
