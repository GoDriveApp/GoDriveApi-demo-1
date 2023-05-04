package Service

import (
	"github.com/GoDriveApp/GoDriveApi/Core/Entity"
	Val "github.com/GoDriveApp/GoDriveApi/Core/Value"
	Repo "github.com/GoDriveApp/GoDriveApi/Domain/Contract/primaryDbRepository"
	"github.com/GoDriveApp/GoDriveApi/Domain/Err"
	"github.com/google/uuid"
)

type (
	AccountFactory struct {
		accRepo         Repo.IAccountRepository
		passwordService IPasswordService
	}

	IAccountFactory interface {
		CreateNewAccount(username Val.Username, email Val.Email, password Val.Password) *Entity.Account
	}
)

func (acf *AccountFactory) CreateNewAccount(username Val.Username, email Val.Email, password Val.Password) (error, *Entity.Account) {
	if acf.accRepo.IsUsernameExist(username) {
		return Err.ThrowDuplicatedUsernameErr(username), nil
	}
	if acf.accRepo.IsEmailExist(email) {
		return Err.ThrowDuplicatedEmailErr(email), nil
	}

	passwordHash := acf.passwordService.Hash(password)
	uid := uuid.New().String()

	return nil, Entity.NewAccount(uid, username, email, passwordHash)
}
