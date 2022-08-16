package validator

import (
	"github.com/go-bp-ikm-api/dto"
	"github.com/go-playground/validator/v10"
)

func ValidateCreateSubCategory(category dto.CreateSubCategoryDto) error {
	var validateCreateSubCategory = validator.New()
	err := validateCreateSubCategory.Struct(category)
	return err
}
