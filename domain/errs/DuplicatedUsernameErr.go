package errs

import (
	"fmt"
	val "github.com/GoDriveApp/GoDriveApi/core/valueObjects"
)

type DuplicatedUsernameErr struct {
	username val.Username
}

func ThrowDuplicatedUsernameErr(username val.Username) DuplicatedUsernameErr {
	return DuplicatedUsernameErr{username}
}

func (err DuplicatedUsernameErr) Error() string {
	return fmt.Sprintf("'%s' is a duplicated username", err.username.GetValue())
}
