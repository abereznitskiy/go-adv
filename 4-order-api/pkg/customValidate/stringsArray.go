package customValidate

import (
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

func StringArrayValidation(fl validator.FieldLevel) bool {
	val, ok := fl.Field().Interface().(pq.StringArray)
	if !ok {
		return false
	}

	for _, v := range val {
		if v == "" {
			return false
		}
	}

	return true
}
