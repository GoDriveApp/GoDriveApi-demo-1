package Err

import "fmt"

type InvalidPasswordHashErr struct {
	passwordHash string
}

func ThrowInvalidPasswordHashErr(passwordHash string) InvalidPasswordHashErr {
	return InvalidPasswordHashErr{passwordHash}
}

func (err InvalidPasswordHashErr) Error() string {
	return fmt.Sprintf("`%s` is an invalid PasswordHash", err.passwordHash)
}
