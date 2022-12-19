package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/my-storage/ms-profile/src/shared/aggregators/errors"
)

func Attest[Object any](object *Object, validate *validator.Validate) (ok bool, err error) {
	if object != nil {
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		if err := validate.Struct(object); err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				return false, err
			}

			for _, err := range err.(validator.ValidationErrors) {
				return false, errors.New("Invalid validation", fmt.Sprintf("Invalid field: '%v' reason: '%v'", err.StructField(), err.Tag()))
			}
		}
	}

	return true, nil
}
