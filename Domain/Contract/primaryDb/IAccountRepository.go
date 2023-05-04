package primaryDb

import (
	"github.com/GoDriveApp/GoDriveApi/Core/Entity"
	"github.com/GoDriveApp/GoDriveApi/Core/Value"
)

type IAccountRepository interface {
	IPrimaryDbRepository[Entity.Account]
	IsEmailExist(email Value.Email) bool
	IsUsernameExist(username Value.Username) bool
}
