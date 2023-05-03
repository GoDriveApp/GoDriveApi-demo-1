package Err

import (
	"fmt"
	"github.com/GoDriveApp/GoDriveApi/Core/Value"
)

type DuplicatedUsernameErr struct {
	username Value.Username
}

func ThrowDuplicatedUsernameErr(username Value.Username) DuplicatedUsernameErr {
	return DuplicatedUsernameErr{username}
}

func (err DuplicatedUsernameErr) Error() string {
	return fmt.Sprintf("'%s' is a duplicated username", err.username.GetValue())
}
