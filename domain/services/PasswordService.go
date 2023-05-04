package services

import (
	val "github.com/GoDriveApp/GoDriveApi/core/valueObjects"
)

type (
	IPasswordService interface {
		Hash(password val.Password) val.PasswordHash
		ComparePassword(password val.Password, hashPassword val.PasswordHash) bool
	}

	PasswordService struct {
	}
)

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (pse PasswordService) Hash(password val.Password) val.PasswordHash {
	//TODO hash password
	_, passwordHash := val.NewPasswordHash(password.GetValue())
	return passwordHash
}

func (pse PasswordService) ComparePassword(password val.Password, hashPassword val.PasswordHash) bool {
	// TODO compare password
	return true
}
