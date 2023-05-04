package account

import (
	"database/sql/driver"
	"errors"
	"github.com/GoDriveApp/GoDriveApi/core/errs"
	"regexp"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type Email struct {
	value string
}

func NewEmail(value string) (error, Email) {
	if !isEmailValueValid(value) {
		return errs.ThrowInvalidEmailErr(value), Email{}
	}
	return nil, Email{value}
}

func isEmailValueValid(value string) bool {
	return emailRegex.MatchString(value)
}

func (eml Email) GetValue() string {
	return eml.value
}

func (eml *Email) Scan(value interface{}) error {
	if str, ok := value.(string); ok {
		eml.value = str
		return nil
	}
	return errors.New("failed to scan Email from database")
}

func (eml Email) Value() (driver.Value, error) {
	return eml.value, nil
}
