package model

type Brand struct {
	Base
	BrandAPI
}

type BrandAPI struct {
	BrandCode *string `json:"brand_code,omitempty" validate:"required"`
	BrandName *string `json:"brand_name,omitempty" validate:"required"`
}
