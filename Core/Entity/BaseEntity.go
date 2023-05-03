package Entity

import "time"

type BaseEntity struct {
	id         string
	insertedAt time.Time
	updatedAt  time.Time
}

func NewBaseEntity(id string) *BaseEntity {
	return &BaseEntity{
		id: id,
	}
}

func (be *BaseEntity) GetId() string {
	return be.id
}

func (be *BaseEntity) GetInsertedAt() time.Time {
	return be.insertedAt
}

func (be *BaseEntity) SetInsertedAt(insertedAt time.Time) {
	be.insertedAt = insertedAt
}

func (be *BaseEntity) GetUpdatedAt() time.Time {
	return be.updatedAt
}

func (be *BaseEntity) SetUpdatedAt(updatedAt time.Time) {
	be.updatedAt = updatedAt
}
