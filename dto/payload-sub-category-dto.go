package dto

type FilterSubCategoryDto struct {
	Name   string `json:"_,validate:"required"`
	Status string `json:"_,validate:"required"`
}

type PageSubCategoryDto struct {
	Name   string `json:"_,validate:"required"`
	Status string `json:"_,validate:"required"`
}

type CreateSubCategoryDto struct {
	Name       string `json:"_,validate:"required"`
	CategoryId string `json:"category_id,validate:"required"`
	Status     string `validate:"required"`
}

type UpdateSubCategoryDto struct {
	Name       string `json:"_,validate:"required"`
	CategoryId string `json:"category_id,validate:"required"`
	Status     string `json:"_,validate:"required"`
}
