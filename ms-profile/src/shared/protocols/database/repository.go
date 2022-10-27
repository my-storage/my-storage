package database

import "github.com/my-storage/ms-profile/src/shared/models"

type ModelExtended interface {
	models.BaseModel
}

type IRepository[Model ModelExtended] interface {
	Add(data Model)
	Update(where Model, data Model)
	Remove(where Model)
	FindOne(where Model)
	FindMany(where Model)
}
