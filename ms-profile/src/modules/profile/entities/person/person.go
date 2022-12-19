package person

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/my-storage/ms-profile/src/modules/authorization/entities/role"
	validatorHelper "github.com/my-storage/ms-profile/src/shared/aggregators/validator"
	"github.com/my-storage/ms-profile/src/shared/utils/time"
)

type Person struct {
	Id        string      `json:"id" validate:"required,uuid4"`
	Name      string      `json:"name" validate:"required,max=100"`
	Email     string      `json:"email" validate:"required,email,max=100"`
	Password  string      `json:"-" validate:"required,max=100"`
	IsActive  bool        `json:"is_active" validate:"required"`
	Roles     []role.Role `json:"roles"`
	CreatedAt int
	UpdatedAt int
	DeletedAt int
}

var validate *validator.Validate

func New(name string, email string, password string) (*Person, error) {
	validate = validator.New()

	id := uuid.New()

	person := &Person{
		Id:        id.String(),
		Name:      name,
		Password:  password,
		Email:     email,
		IsActive:  true,
		CreatedAt: time.GetCurrentTime(),
	}

	if err := person.Validate(); err != nil {
		return nil, err
	}

	return person, nil
}

func (p *Person) Validate() error {
	_, err := validatorHelper.Attest(p, validate)

	return err
}

func (p *Person) AddRole(roles ...role.Role) {
	p.Roles = append(p.Roles, roles...)
}

func (p *Person) Activate() {
	p.IsActive = true
}

func (p *Person) Deactivate() {
	p.IsActive = false
}
