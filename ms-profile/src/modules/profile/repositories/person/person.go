package person

import "github.com/my-storage/ms-profile/src/modules/profile/entities/person"

type Where struct {
	Id        string
	Name      string
	Email     string
	IsActive  bool
	CreatedAt int
	UpdatedAt int
	DeletedAt int
}

type Options struct {
	Page    int
	PerPage int
	Order   string
}

type PersonRepository interface {
	Insert(data person.Person) (insertedCount int, err error)
	Update(Id string, data person.Person) (updatedCount int, err error)
	Remove(Id string) (removedCount int, err error)
	FindOne(where Where) (person *person.Person, err error)
	FindMany(where Where, options Options) (person *person.Person, err error)
}
