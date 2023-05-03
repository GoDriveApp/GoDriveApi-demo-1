package Err

import "fmt"

type InvalidUsernameErr struct {
	username string
}

func ThrowInvalidUsernameErr(username string) InvalidUsernameErr {
	return InvalidUsernameErr{username}
}

func (err InvalidUsernameErr) Error() string {
	return fmt.Sprintf("`%s` is an invalid Username", err.username)
}
