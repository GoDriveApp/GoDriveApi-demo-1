package primaryDbRepository

import "github.com/GoDriveApp/GoDriveApi/Core/Value"

type IAccountRepository interface {
	IPrimaryDbRepository
	IsEmailExist(email Value.Email) bool
	IsUsernameExist(username Value.Username) bool
}
