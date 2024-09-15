package validator

import (
	"github.com/go-playground/validator"
	"reflect"
	"tender_api/src/internal/model/enum"
)

func validateTenderServiceType(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return enum.TenderServiceType(field.String()).IsValid()
	case reflect.Slice:
		for i := 0; i < field.Len(); i++ {
			value, ok := field.Index(i).Interface().(enum.TenderServiceType)
			if !ok || !value.IsValid() {
				return false
			}
		}
		return true
	default:
		return false
	}
}
