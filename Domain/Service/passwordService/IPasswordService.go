package passwordService

import "github.com/GoDriveApp/GoDriveApi/Core/Value"

type IPasswordService interface {
	Hash(password Value.Password) Value.PasswordHash
	ComparePassword(password Value.Password, hashPassword Value.PasswordHash) bool
}
