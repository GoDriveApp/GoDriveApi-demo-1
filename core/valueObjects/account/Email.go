package account

import (
	"github.com/GoDriveApp/GoDriveApi/core/Err"
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
		return Err.ThrowInvalidEmailErr(value), Email{}
	}
	return nil, Email{value}
}

func isEmailValueValid(value string) bool {
	return emailRegex.MatchString(value)
}

func (eml Email) GetValue() string {
	return eml.value
}
