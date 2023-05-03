package Err

import (
	"fmt"
	"github.com/GoDriveApp/GoDriveApi/Core/Value"
)

type DuplicatedEmailErr struct {
	email Value.Email
}

func ThrowDuplicatedEmailErr(email Value.Email) DuplicatedEmailErr {
	return DuplicatedEmailErr{email}
}

func (err DuplicatedEmailErr) Error() string {
	return fmt.Sprintf("'%s' is a duplicated email", err.email.GetValue())
}
