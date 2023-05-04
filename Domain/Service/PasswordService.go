package Service

import "github.com/GoDriveApp/GoDriveApi/Core/Value"

type (
	IPasswordService interface {
		Hash(password Value.Password) Value.PasswordHash
		ComparePassword(password Value.Password, hashPassword Value.PasswordHash) bool
	}

	PasswordService struct {
	}
)

func (pse PasswordService) Hash(password Value.Password) Value.PasswordHash {
	//TODO hash password
	_, passwordHash := Value.NewPasswordHash(password.GetValue())
	return passwordHash
}

func (pse PasswordService) ComparePassword(password Value.Password, hashPassword Value.PasswordHash) bool {
	// TODO compare password
	return true
}
