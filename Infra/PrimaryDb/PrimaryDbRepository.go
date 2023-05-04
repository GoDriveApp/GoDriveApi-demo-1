package PrimaryDb

import "github.com/GoDriveApp/GoDriveApi/Core/Entity"

// TODO implement all functions

type PrimaryDbRepository[T Entity.BaseEntity] struct {
}

func NewPrimaryDbRepository[T Entity.BaseEntity]() PrimaryDbRepository[T] {
	return PrimaryDbRepository[T]{}
}

func (repo *PrimaryDbRepository[T]) Save(entity T) {

}

func (repo *PrimaryDbRepository[T]) Update(entity T) {

}

func (repo *PrimaryDbRepository[T]) Delete(entity T) {

}

func (repo *PrimaryDbRepository[T]) GetById(id string) *T {
	return nil
}
