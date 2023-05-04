package valueObjects

import "github.com/GoDriveApp/GoDriveApi/core/valueObjects/account"

// for account
type (
	Username     = account.Username
	Email        = account.Email
	PasswordHash = account.PasswordHash
	Password     = account.Password
)

var (
	NewUsername     = account.NewUsername
	NewEmail        = account.NewEmail
	NewPasswordHash = account.NewPasswordHash
	NewPassword     = account.NewPassword
)
