package main

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func ValidatePaste(paste *PasteRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(paste)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
