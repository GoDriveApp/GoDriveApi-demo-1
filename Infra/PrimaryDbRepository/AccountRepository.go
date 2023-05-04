package PrimaryDbRepository

import (
	"github.com/GoDriveApp/GoDriveApi/Core/Entity"
	"github.com/GoDriveApp/GoDriveApi/Core/Value"
)

type AccountRepository struct {
	PrimaryDbRepository[Entity.Account]
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (accRepo *AccountRepository) IsEmailExist(email Value.Email) bool {
	return true
}

func (accRepo *AccountRepository) IsUsernameExist(username Value.Username) bool {
	return true
}
