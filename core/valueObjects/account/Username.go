package account

import (
	"database/sql/driver"
	"errors"
	"github.com/GoDriveApp/GoDriveApi/core/errs"
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
		return errs.ThrowInvalidUsernameErr(value), Username{}
	}
	return nil, Username{value}
}

func isUsernameValueValid(value string) bool {
	return usernameRegex.MatchString(value)
}

func (usn Username) GetValue() string {
	return usn.value
}

func (usn *Username) Scan(value any) error {
	if str, ok := value.(string); ok {
		usn.value = str
		return nil
	}
	return errors.New("failed to scan Username from database")
}

func (usn Username) Value() (driver.Value, error) {
	return usn.value, nil
}
