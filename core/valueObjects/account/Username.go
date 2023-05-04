package account

import (
	"github.com/GoDriveApp/GoDriveApi/core/Err"
	"regexp"
)

var (
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{1,30}$`)
)

type Username struct {
	value string
}

func NewUsername(value string) (error, Username) {
	if !isUsernameValueValid(value) {
		return Err.ThrowInvalidUsernameErr(value), Username{}
	}
	return nil, Username{value}
}

func isUsernameValueValid(value string) bool {
	return usernameRegex.MatchString(value)
}

func (usn Username) GetValue() string {
	return usn.value
}
