package repositories

import ent "github.com/GoDriveApp/GoDriveApi/core/entities"

type BaseRepository[T ent.BaseEntity] struct {
}

func (r *BaseRepository[T]) Save(entity T) {
	//TODO implement me
	panic("implement me")
}

func (r *BaseRepository[T]) Update(entity T) {
	//TODO implement me
	panic("implement me")
}

func (r *BaseRepository[T]) Delete(entity T) {
	//TODO implement me
	panic("implement me")
}

func (r *BaseRepository[T]) GetById(id string) *T {
	//TODO implement me
	panic("implement me")
}
