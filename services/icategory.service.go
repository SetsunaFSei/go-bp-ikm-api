package services

import (
	"github.com/go-bp-ikm-api/dto"
	"github.com/go-bp-ikm-api/mapping"
)

type ICategoryService interface {
	CreateCategory(*dto.CreateCategoryDto) error
	GetAll() ([]*mapping.CategoryVm, error)
	GetCategoryById(*string) *mapping.CategoryVm

	// UpdateCategory(*dto.FilterCategoryDto) error
	// DeleteCategory(*string) error
}
