package permission

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/my-storage/ms-profile/src/shared/aggregators/errors"
)

func TestNew(t *testing.T) {
	t.Run("Should return a new permission correctly", func(t *testing.T) {
		name := "permission name"
		slug := "permission-name"
		description := "permission description"

		permission, err := New(name, slug, description)

		assert.Nil(t, err)
		assert.Equal(t, permission.Name, name)
		assert.Equal(t, permission.Description, description)
		assert.Equal(t, permission.Slug, slug)
	})

	t.Run("Should return a error because 'Name' is required", func(t *testing.T) {
		permission, err := New("", "permission-name", "permission description")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Name' reason: 'required'"), err)
		assert.Nil(t, permission)
	})

	t.Run("Should return a error because 'Slug' is required", func(t *testing.T) {
		permission, err := New("permission name", "", "permission description")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Slug' reason: 'required'"), err)
		assert.Nil(t, permission)
	})

	t.Run("Should return a error because 'Name' have max length in 100", func(t *testing.T) {
		permission, err := New("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut abore et", "permission-name", "permission description")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Name' reason: 'max'"), err)
		assert.Nil(t, permission)
	})

	t.Run("Should return a error because 'Slug' have max length in 100", func(t *testing.T) {
		permission, err := New("permission name", "lorem_ipsum_dolor_sit_amet_consectetur_adipiscing_elit_sed_do_eiusmod_temporsdf_incididunt_ut_abore_et", "permission description")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Slug' reason: 'max'"), err)
		assert.Nil(t, permission)
	})

	t.Run("Should return a error because 'Description' have max length in 250", func(t *testing.T) {
		permission, err := New(
			"permission name",
			"permission-name",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut abore et Lorem ipsum dolor sit amet, consectetur adipiscing, elit, sed do eiusmod tempor incididunt ut abore, et Lorem ipsum dolor sit, consectetur adipiscing",
		)
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Description' reason: 'max'"), err)
		assert.Nil(t, permission)
	})
}
