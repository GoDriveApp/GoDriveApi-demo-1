package account

import (
	"github.com/GoDriveApp/GoDriveApi/Core/Err"
	"regexp"
)

var (
	passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{8,}$`)
)

type Password struct {
	value string
}

func NewPassword(value string) (error, Password) {
	if !isPasswordValueValid(value) {
		return Err.ThrowInvalidPasswordErr(value), Password{}
	}
	return nil, Password{value}
}

func isPasswordValueValid(value string) bool {
	return passwordRegex.MatchString(value)
}
