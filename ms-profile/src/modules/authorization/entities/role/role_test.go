package role

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/my-storage/ms-profile/src/modules/authorization/entities/permission"
	"github.com/my-storage/ms-profile/src/shared/aggregators/errors"
)

func TestNew(t *testing.T) {
	t.Run("Should return a new role correctly", func(t *testing.T) {
		name := "role name"
		slug := "role-name"
		description := "role description"

		role, err := New(name, slug, description)

		assert.Nil(t, err)
		assert.Equal(t, role.Name, name)
		assert.Equal(t, role.Description, description)
		assert.Equal(t, role.Slug, slug)
	})

	t.Run("Should return a error because 'Name' is required", func(t *testing.T) {
		role, err := New("", "role-name", "role description")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Name' reason: 'required'"), err)
		assert.Nil(t, role)
	})

	t.Run("Should return a error because 'Slug' is required", func(t *testing.T) {
		role, err := New("role name", "", "role description")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Slug' reason: 'required'"), err)
		assert.Nil(t, role)
	})

	t.Run("Should return a error because 'Name' have max length in 100", func(t *testing.T) {
		role, err := New("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut abore et", "role-name", "role description")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Name' reason: 'max'"), err)
		assert.Nil(t, role)
	})

	t.Run("Should return a error because 'Slug' have max length in 100", func(t *testing.T) {
		role, err := New("role name", "lorem_ipsum_dolor_sit_amet_consectetur_adipiscing_elit_sed_do_eiusmod_temporsdf_incididunt_ut_abore_et", "role description")
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Slug' reason: 'max'"), err)
		assert.Nil(t, role)
	})

	t.Run("Should return a error because 'Description' have max length in 250", func(t *testing.T) {
		role, err := New(
			"role name",
			"role-name",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut abore et Lorem ipsum dolor sit amet, consectetur adipiscing, elit, sed do eiusmod tempor incididunt ut abore, et Lorem ipsum dolor sit, consectetur adipiscing",
		)
		assert.Equal(t, errors.New("Invalid validation", "Invalid field: 'Description' reason: 'max'"), err)
		assert.Nil(t, role)
	})
}

func TestAttach(t *testing.T) {
	t.Run("Should test a Attach method correctly", func(t *testing.T) {
		role, err := New("role name", "role-name", "role description")
		assert.Nil(t, err)

		perm, err := permission.New("permission name", "permission-name", "permission description")
		assert.Nil(t, err)

		role.Attach(*perm)
		assert.Equal(t, role.Permissions, []permission.Permission{*perm})
	})
}
