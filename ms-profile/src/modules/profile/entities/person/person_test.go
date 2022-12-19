package person

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/my-storage/ms-profile/src/modules/authorization/entities/role"
	"github.com/my-storage/ms-profile/src/shared/aggregators/errors"
)

func TestNew(t *testing.T) {
	t.Run("Should return a new person correctly", func(t *testing.T) {
		name := "John Doe"
		email := "example@email.com"
		password := "password"

		person, err := New(name, email, password)

		assert.Nil(t, err)
		assert.Equal(t, person.Email, email)
		assert.Equal(t, person.Name, name)
		assert.Equal(t, person.Password, password)
	})

	t.Run("Should return a new person with 'IsAction' on 'true'", func(t *testing.T) {
		person, err := New("John Doe", "example@email.com", "password")
		assert.Nil(t, err)
		assert.NotNil(t, person)
		assert.True(t, person.IsActive)
	})

	t.Run("Should return a error because 'Name' is required", func(t *testing.T) {
		person, err := New("", "example@email.com", "password")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Name' reason: 'required'"), err)
		assert.Nil(t, person)
	})

	t.Run("Should return a error because 'Name' have max length in 100", func(t *testing.T) {
		person, err := New("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut abore et", "example@email.com", "password")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Name' reason: 'max'"), err)
		assert.Nil(t, person)
	})

	t.Run("Should return a error because 'Email' have max length in 100", func(t *testing.T) {
		person, err := New("John Dow", "ldfforemipsumdolorsitasdfqwmetconsecteturadipiscingelitseddoeiusmodtemporincididuntutaboret@email.com", "password")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Email' reason: 'max'"), err)
		assert.Nil(t, person)
	})

	t.Run("Should return a error because 'Email' is invalid", func(t *testing.T) {
		person, err := New("John Doe", "non-valid-email", "password")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Email' reason: 'email'"), err)
		assert.Nil(t, person)
	})

	t.Run("Should return a error because 'Email' is required", func(t *testing.T) {
		person, err := New("John Doe", "", "password")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Email' reason: 'required'"), err)
		assert.Nil(t, person)
	})

	t.Run("Should return a error because 'Password' is required", func(t *testing.T) {
		person, err := New("John Doe", "example@email.com", "")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Password' reason: 'required'"), err)
		assert.Nil(t, person)
	})

	t.Run("Should return a error because 'Password' have max length in 100", func(t *testing.T) {
		person, err := New("John Dow", "example@email.com", "ldfforemipsusdmdfdolsdforsitasdfqwmetconsecteturadipiscingelitseddoeiusmodtemporincididuntutaborfdset")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Password' reason: 'max'"), err)
		assert.Nil(t, person)
	})
}

func TestActivate(t *testing.T) {
	t.Run("Should test a Activate method correctly", func(t *testing.T) {
		person, err := New("John Doe", "john@example.com", "password")
		assert.Nilf(t, err, "Unexpected error received: %v", err)

		person.Activate()
		assert.True(t, person.IsActive)
	})
}

func TestDeactivate(t *testing.T) {
	t.Run("Should test a Deactivate method correctly", func(t *testing.T) {
		person, err := New("John Doe", "john@example.com", "password")
		assert.Nilf(t, err, "Unexpected error received: %v", err)

		person.Deactivate()
		assert.False(t, person.IsActive)
	})
}

func TestAddRole(t *testing.T) {
	t.Run("Should test a AddRole method correctly", func(t *testing.T) {
		person, err := New("John Doe", "john@example.com", "password")
		assert.Nilf(t, err, "Unexpected error received: %v", err)

		rol, err := role.New("role name", "role-name", "role description")
		assert.Nil(t, err)

		person.AddRole(*rol)
		assert.Equal(t, person.Roles, []role.Role{*rol})
	})
}
