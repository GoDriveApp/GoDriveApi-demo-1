package passwordService

import "github.com/GoDriveApp/GoDriveApi/Core/Value"

type PasswordService struct {
}

func (pse PasswordService) Hash(password Value.Password) Value.PasswordHash {
	//TODO hash password
	_, passwordHash := Value.NewPasswordHash(password.GetValue())
	return passwordHash
}

func (pse PasswordService) ComparePassword(password Value.Password, hashPassword Value.PasswordHash) bool {
	// TODO compare password
	return true
}
