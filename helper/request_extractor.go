package helper

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ExtractValidationErrors(err error, obj interface{}) map[string]string {
	errors := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, v := range validationErrors {
			// Get JSON field name
			field, _ := reflect.TypeOf(obj).FieldByName(v.Field())
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				jsonTag = v.Field() // Fallback to struct field name if no json tag
			}

			errors[jsonTag] = "Field validation for '" + jsonTag + "' failed on the '" + v.Tag() + "' tag"
		}
	}
	return errors
}
