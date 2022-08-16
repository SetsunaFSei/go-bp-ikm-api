package validator

import (
	"github.com/go-bp-ikm-api/dto"
	"github.com/go-playground/validator/v10"
)

func ValidateCreateCategory(category dto.CreateCategoryDto) error {
	var validateCreateCategory = validator.New()
	err := validateCreateCategory.Struct(category)
	return err
}
