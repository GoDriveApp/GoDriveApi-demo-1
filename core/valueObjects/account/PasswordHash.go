package account

import (
	"github.com/GoDriveApp/GoDriveApi/core/errs"
	"regexp"
)

var (
	passwordHashRegex = regexp.MustCompile(`^.{51,}$`)
)

type PasswordHash struct {
	value string
}

func NewPasswordHash(value string) (error, PasswordHash) {
	if !isPasswordHashValueValid(value) {
		return errs.ThrowInvalidPasswordHashErr(value), PasswordHash{}
	}
	return nil, PasswordHash{value}
}

func isPasswordHashValueValid(value string) bool {
	return passwordHashRegex.MatchString(value)
}

func (psh PasswordHash) GetValue() string {
	return psh.value
}
