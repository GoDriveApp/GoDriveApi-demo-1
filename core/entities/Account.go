package entities

import (
	val "github.com/GoDriveApp/GoDriveApi/core/valueObjects"
)

type Account struct {
	BaseEntity
	username     val.Username `gorm:"uniqueIndex"`
	email        val.Email    `gorm:"uniqueIndex"`
	passwordHash val.PasswordHash
}

func NewAccount(id string, username val.Username, email val.Email, passwordHash val.PasswordHash) *Account {
	return &Account{
		BaseEntity:   NewBaseEntity(id),
		username:     username,
		email:        email,
		passwordHash: passwordHash,
	}
}

func (acc *Account) GetUsername() val.Username {
	return acc.username
}

func (acc *Account) GetEmail() val.Email {
	return acc.email
}

func (acc *Account) GetPasswordHash() val.PasswordHash {
	return acc.passwordHash
}

func (acc *Account) SetPasswordHash(passwordHash val.PasswordHash) {
	acc.passwordHash = passwordHash
}
