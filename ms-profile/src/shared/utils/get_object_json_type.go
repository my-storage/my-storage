package utils

import (
	"errors"
	"reflect"
)

func GetObjectJSONType(obj reflect.Type) (*string, error) {
	var result string

	switch obj.Kind() {
	case reflect.Int:
		result = "number"
	case reflect.Int8:
		result = "number"
	case reflect.Int16:
		result = "number"
	case reflect.Int32:
		result = "number"
	case reflect.Int64:
		result = "number"
	case reflect.Uint:
		result = "number"
	case reflect.Uint8:
		result = "number"
	case reflect.Uint16:
		result = "number"
	case reflect.Uint32:
		result = "number"
	case reflect.Uint64:
		result = "number"
	case reflect.Uintptr:
		result = "number"
	case reflect.Float32:
		result = "number"
	case reflect.Float64:
		result = "number"
	case reflect.Complex64:
		result = "number"
	case reflect.Complex128:
		result = "number"
	case reflect.String:
		result = "string"
	case reflect.Bool:
		result = "boolean"
	case reflect.Struct:
		result = "object"
	case reflect.Map:
		result = "object"
	case reflect.Array:
		result = "array"
	case reflect.Slice:
		result = "array"

	default:
		return nil, errors.New("unsupported reflect.Type to convert to JSON object type")
	}

	return &result, nil
}
