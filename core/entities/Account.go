package entities

import (
	val "github.com/GoDriveApp/GoDriveApi/core/valueObjects"
)

type Account struct {
	BaseEntity
	Username     val.Username     `gorm:"uniqueIndex;type:varchar(30)"`
	Email        val.Email        `gorm:"uniqueIndextype:varchar"`
	PasswordHash val.PasswordHash `gorm:"varchar"`
}

func NewAccount(id string, username val.Username, email val.Email, passwordHash val.PasswordHash) *Account {
	return &Account{
		BaseEntity:   NewBaseEntity(id),
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}
}

func (acc *Account) GetUsername() val.Username {
	return acc.Username
}

func (acc *Account) GetEmail() val.Email {
	return acc.Email
}

func (acc *Account) GetPasswordHash() val.PasswordHash {
	return acc.PasswordHash
}

func (acc *Account) SetPasswordHash(passwordHash val.PasswordHash) {
	acc.PasswordHash = passwordHash
}
