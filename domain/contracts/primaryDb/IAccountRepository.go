package primaryDb

import (
	ent "github.com/GoDriveApp/GoDriveApi/core/entities"
	val "github.com/GoDriveApp/GoDriveApi/core/valueObjects"
)

type IAccountRepository interface {
	IRepository[ent.Account]
	IsEmailExist(email val.Email) bool
	IsUsernameExist(username val.Username) bool
}
