package account

import (
	"github.com/GoDriveApp/GoDriveApi/core/errs"
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
		return errs.ThrowInvalidPasswordErr(value), Password{}
	}
	return nil, Password{value}
}

func isPasswordValueValid(value string) bool {
	return passwordRegex.MatchString(value)
}

func (pss Password) GetValue() string {
	return pss.value
}
