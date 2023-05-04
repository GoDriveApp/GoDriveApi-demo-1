package errs

import (
	"fmt"
	val "github.com/GoDriveApp/GoDriveApi/core/valueObjects"
)

type DuplicatedEmailErr struct {
	email val.Email
}

func ThrowDuplicatedEmailErr(email val.Email) DuplicatedEmailErr {
	return DuplicatedEmailErr{email}
}

func (err DuplicatedEmailErr) Error() string {
	return fmt.Sprintf("'%s' is a duplicated email", err.email.GetValue())
}
