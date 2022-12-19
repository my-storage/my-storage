package permission

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	validatorHelper "github.com/my-storage/ms-profile/src/shared/aggregators/validator"
)

type Permission struct {
	Id          string `json:"id" validate:"required,uuid4"`
	Name        string `json:"name" validate:"required,max=100"`
	Slug        string `json:"slug" validate:"required,max=100"`
	Description string `json:"description" validate:"omitempty,max=250"`
}

var validate *validator.Validate

func New(name string, slug string, description string) (*Permission, error) {
	validate = validator.New()

	id := uuid.New()

	perm := &Permission{
		Id:          id.String(),
		Name:        name,
		Slug:        slug,
		Description: description,
	}

	if err := perm.Validate(); err != nil {
		return nil, err
	}

	return perm, nil
}

func (p *Permission) Validate() error {
	_, err := validatorHelper.Attest(p, validate)

	return err
}
