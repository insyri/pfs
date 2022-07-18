package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/insyri/pfs/backend/structures"
)

var validate = validator.New()

func ValidatePaste(paste *structures.PasteRequest) []*structures.ErrorResponse {
	var errors []*structures.ErrorResponse
	err := validate.Struct(paste)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element structures.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
