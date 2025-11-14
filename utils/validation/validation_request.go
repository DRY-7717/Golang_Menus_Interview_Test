package validation

import "github.com/go-playground/validator/v10"

func CustomValidator(errStruct error) []interface{} {
	var errors []interface{}

	validationErrors := errStruct.(validator.ValidationErrors)

	for _, error := range validationErrors {
		errors = append(errors, "The field "+error.Field()+" is "+error.Tag()+" "+error.Param())
	}

	return errors
}
