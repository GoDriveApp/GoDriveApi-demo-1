package accountFactory

import (
	"github.com/GoDriveApp/GoDriveApi/Core/Entity"
	"github.com/GoDriveApp/GoDriveApi/Core/Value"
)

type IAccountFactory interface {
	CreateNewAccount(username Value.Username, email Value.Email, password Value.Password) *Entity.Account
}
