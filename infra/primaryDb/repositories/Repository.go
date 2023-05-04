package repositories

import ent "github.com/GoDriveApp/GoDriveApi/core/entities"

type Repository[T ent.BaseEntity] struct {
}

func (r *Repository[T]) Save(entity T) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository[T]) Update(entity T) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository[T]) Delete(entity T) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository[T]) GetById(id string) *T {
	//TODO implement me
	panic("implement me")
}
