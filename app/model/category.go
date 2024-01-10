package model

type Category struct {
	Base
	CategoryAPI
}

type CategoryAPI struct {
	CategoryCode *string `json:"category_code,omitempty" validate:"required"`
	CategoryName *string `json:"category_name,omitempty" validate:"required"`
}
