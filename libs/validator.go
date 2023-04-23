package libs

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InputValidator(payload any) string {
	var message string

	// Instantiate validator package
	validate = validator.New()

	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				message = err.StructField() + " field is required"
				break
			}

			if err.Tag() == "oneof" {
				message = err.StructField() + " needs to be 0 or 1, 0 for male, 1 for female"
				break
			}

			if err.Tag() == "gte" {
				message = err.StructField() + " needs to be higher or equal than 0"
				break
			}

			if err.Tag() == "lte" {
				message = err.StructField() + " needs to be lower or equal than 120"
				break
			}

			message = err.Error()
		}
	}

	return message
}
