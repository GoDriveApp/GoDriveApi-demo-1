package Err

import "fmt"

type InvalidEmailErr struct {
	email string
}

func ThrowInvalidEmailErr(email string) InvalidEmailErr {
	return InvalidEmailErr{email}
}

func (err InvalidEmailErr) Error() string {
	return fmt.Sprintf("`%s` is an invalid Email", err.email)
}
