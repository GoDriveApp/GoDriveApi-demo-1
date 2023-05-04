package primaryDb

import (
	ent "github.com/GoDriveApp/GoDriveApi/core/entities"
)

type IRepository[T ent.BaseEntity] interface {
	Save(entity T)
	Update(entity T)
	Delete(entity T)
	GetById(id string) *T
}
