package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/my-storage/ms-profile/src/shared/aggregators/errors"
	"github.com/my-storage/ms-profile/src/shared/utils"
)

func GetBody[Body any](c *gin.Context, validate *validator.Validate) *Body {
	body := new(Body)

	decodeBody(body, c.Request.Body)

	validateBody(body, validate)

	return body
}

func SetHeaders(c *gin.Context, headers http.Header) {
	for key, values := range headers {
		parsedValues := strings.Join(values[:], ",")

		c.Writer.Header().Set(key, parsedValues)
	}
}

func decodeBody(target any, source io.Reader) {
	if source != nil {
		decoder := json.NewDecoder(source)

		result := make(map[string]string)

		if err := decoder.Decode(target); err != nil && err != io.EOF {
			if _, ok := err.(*http.MaxBytesError); ok {
				panic(errors.New(errors.RequestEntityTooLarge, "The request's size exceeds the server's size limit.", nil))
			}

			error := err.(*json.UnmarshalTypeError)

			expectedType, err := utils.GetObjectJSONType(error.Type)
			if err != nil {
				panic(errors.New(errors.BadRequest, err.Error(), nil))
			}

			message := fmt.Sprintf("Invalid type on field '%v', expected type '%v' but received '%v'", error.Field, *expectedType, error.Value)

			result[error.Field] = message

			panic(errors.New(errors.BadRequest, "Invalid fields", map[string]any{
				"invalid": result,
			}))
		}
	}
}

func validateBody(obj any, validate *validator.Validate) {
	if obj != nil {
		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		if err := validate.Struct(obj); err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				panic(err)
			}

			result := make(map[string]string)

			for _, err := range err.(validator.ValidationErrors) {
				namespaceSlice := strings.Split(err.Namespace(), ".")
				keySlice := append(namespaceSlice[:0], namespaceSlice[0+1:]...)
				key := strings.Join(keySlice[:], ".")
				expectation := err.Tag()

				var message string

				if expectation == "required" {
					message = fmt.Sprintf("Field '%v' is required", key)
				} else {
					receivedType := reflect.TypeOf(err.Value())
					message = fmt.Sprintf("Invalid type on field '%v', expected type '%v' but received '%v'", key, expectation, receivedType)
				}

				result[key] = message
			}

			panic(errors.New(errors.BadRequest, "Invalid fields", map[string]any{
				"invalid": result,
			}))

		}
	}
}
