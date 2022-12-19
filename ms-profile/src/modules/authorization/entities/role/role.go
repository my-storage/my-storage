package role

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/my-storage/ms-profile/src/modules/authorization/entities/permission"
	validatorHelper "github.com/my-storage/ms-profile/src/shared/aggregators/validator"
)

type Role struct {
	Id          string                  `json:"id" validate:"required,uuid4"`
	Name        string                  `json:"name" validate:"required,max=100"`
	Slug        string                  `json:"slug" validate:"required,max=100"`
	Description string                  `json:"description" validate:"omitempty,max=250"`
	Permissions []permission.Permission `json:"permissions" validate:"required"`
}

var validate *validator.Validate

func New(name string, slug string, description string) (*Role, error) {
	validate = validator.New()

	id := uuid.New()

	role := &Role{
		Id:          id.String(),
		Name:        name,
		Slug:        slug,
		Description: description,
		Permissions: []permission.Permission{},
	}

	if err := role.Validate(); err != nil {
		return nil, err
	}

	return role, nil
}

func (r *Role) Attach(permissions ...permission.Permission) {
	r.Permissions = append(r.Permissions, permissions...)
}

func (r *Role) Validate() error {
	_, err := validatorHelper.Attest(r, validate)

	return err
}
