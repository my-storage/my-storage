package errors

import (
	"fmt"
	"net/http"
)

type ErrorName string

const (
	UnprocessableEntity ErrorName = "UnprocessableEntity"
	NotFound            ErrorName = "NotFound"
	Forbidden           ErrorName = "Forbidden"
	Unauthorized        ErrorName = "Unauthorized"
	BadRequest          ErrorName = "BadRequest"
	InternalServerError ErrorName = "InternalServerError"
)

type AppError struct {
	Description string
	Name        ErrorName
	Details     *any
}

func New(name ErrorName, description string, details any) error {
	return &AppError{
		Description: description,
		Name:        name,
		Details:     &details,
	}
}

func (e AppError) Error() string {
	return fmt.Sprintf("AppError: %v", e.Description)
}

func (e *AppError) GetStatusCode() *int {
	var code int

	switch e.Name {
	case UnprocessableEntity:
		code = http.StatusUnprocessableEntity
	case NotFound:
		code = http.StatusNotFound
	case Forbidden:
		code = http.StatusForbidden
	case Unauthorized:
		code = http.StatusUnauthorized
	case BadRequest:
		code = http.StatusBadRequest
	case InternalServerError:
		code = http.StatusInternalServerError
	default:
		return nil
	}

	return &code
}
