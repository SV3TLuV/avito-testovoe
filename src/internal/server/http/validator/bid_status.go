package validator

import (
	"github.com/go-playground/validator"
	"reflect"
	"tender_api/src/internal/model/enum"
)

func validateBidStatus(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return enum.BidStatus(field.String()).IsValid()
	case reflect.Slice:
		for i := 0; i < field.Len(); i++ {
			value, ok := field.Index(i).Interface().(enum.BidStatus)
			if !ok || !value.IsValid() {
				return false
			}
		}
		return true
	default:
		return false
	}
}
