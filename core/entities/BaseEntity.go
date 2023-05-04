package entities

import "time"

type BaseEntity struct {
	Id         string    `gorm:"primaryKey"`
	InsertedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func NewBaseEntity(id string) BaseEntity {
	return BaseEntity{
		Id: id,
	}
}

func (be *BaseEntity) GetId() string {
	return be.Id
}

func (be *BaseEntity) GetInsertedAt() time.Time {
	return be.InsertedAt
}

func (be *BaseEntity) SetInsertedAt(insertedAt time.Time) {
	be.InsertedAt = insertedAt
}

func (be *BaseEntity) GetUpdatedAt() time.Time {
	return be.UpdatedAt
}

func (be *BaseEntity) SetUpdatedAt(updatedAt time.Time) {
	be.UpdatedAt = updatedAt
}
