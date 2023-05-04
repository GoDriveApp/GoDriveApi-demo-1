package primaryDbRepository

import "github.com/GoDriveApp/GoDriveApi/Core/Entity"

type IPrimaryDbRepository[T Entity.BaseEntity] interface {
	Save(entity T)
	Update(entity T)
	Delete(entity T)
	GetById(id string) *T
}
