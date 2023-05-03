package Entity

import (
	Val "github.com/GoDriveApp/GoDriveApi/Core/Value"
)

type Account struct {
	username     Val.Username
	email        Val.Email
	passwordHash Val.PasswordHash
}

func NewAccount(username Val.Username, email Val.Email, passwordHash Val.PasswordHash) *Account {
	return &Account{username, email, passwordHash}
}

func (acc *Account) GetUsername() Val.Username {
	return acc.username
}

func (acc *Account) GetEmail() Val.Email {
	return acc.email
}

func (acc *Account) GetPasswordHash() Val.PasswordHash {
	return acc.passwordHash
}

func (acc *Account) SetPasswordHash(passwordHash Val.PasswordHash) {
	acc.passwordHash = passwordHash
}
