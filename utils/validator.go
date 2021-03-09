package utils

import (
	"github.com/go-playground/validator/v10"
)

// ErrorResponse validation resp
type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

// ValidateStruct validate req
func ValidateStruct(req interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var e ErrorResponse
			e.FailedField = err.Field()
			e.Tag = err.Tag()
			e.Value = err.Param()
			errors = append(errors, &e)
		}
	}

	return errors
}
