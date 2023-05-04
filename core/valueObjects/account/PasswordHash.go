package account

import (
	"database/sql/driver"
	"errors"
	"regexp"
)

var (
	passwordHashRegex = regexp.MustCompile(`^.{51,}$`)
)

type PasswordHash struct {
	value string
}

func NewPasswordHash(value string) (error, PasswordHash) {
	//if !isPasswordHashValueValid(value) {
	//	return errs.ThrowInvalidPasswordHashErr(value), PasswordHash{}
	//}
	return nil, PasswordHash{value}
}

func isPasswordHashValueValid(value string) bool {
	return passwordHashRegex.MatchString(value)
}

func (psh PasswordHash) GetValue() string {
	return psh.value
}

func (psh *PasswordHash) Scan(value interface{}) error {
	if str, ok := value.(string); ok {
		psh.value = str
		return nil
	}
	return errors.New("failed to scan PasswordHash from database")
}

func (psh PasswordHash) Value() (driver.Value, error) {
	return psh.value, nil
}
