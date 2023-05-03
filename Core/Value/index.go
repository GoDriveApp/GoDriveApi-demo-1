package Value

import "github.com/GoDriveApp/GoDriveApi/Core/Value/account"

// for Account
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
