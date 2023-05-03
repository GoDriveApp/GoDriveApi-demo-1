package Err

import "fmt"

type InvalidPasswordErr struct {
	password string
}

func ThrowInvalidPasswordErr(password string) InvalidPasswordErr {
	return InvalidPasswordErr{password}
}

func (err InvalidPasswordErr) Error() string {
	return fmt.Sprintf("`%s` is an invalid Password", err.password)
}
