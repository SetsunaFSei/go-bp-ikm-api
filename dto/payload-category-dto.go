package dto

type FilterCategoryDto struct {
	Name   string `json:"name,validate:"required"`
	Status string `json:"status,validate:"required"`
}

type PageCategoryDto struct {
	Name   string `json:"name,validate:"required"`
	Status string `json:"status,validate:"required"`
}

type CreateCategoryDto struct {
	Id     string `json:"id,validate:"required"`
	Name   string `json:"name,validate:"required"`
	Status string `json:"status,validate:"required"`
}

type UpdateCategoryDto struct {
	Name   string `json:"name,validate:"required"`
	Status string `json:"status,validate:"required"`
}
