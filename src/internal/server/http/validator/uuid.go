package validator

import (
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"reflect"
)

func validateUUID(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		_, err := uuid.Parse(field.String())
		return err == nil
	case reflect.Slice:
		for i := 0; i < field.Len(); i++ {
			if _, err := uuid.Parse(field.Index(i).String()); err != nil {
				return false
			}
		}
		return true
	default:
		return false
	}
}
