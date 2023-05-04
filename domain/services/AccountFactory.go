package services

import (
	"github.com/GoDriveApp/GoDriveApi/core/entities"
	val "github.com/GoDriveApp/GoDriveApi/core/valueObjects"
	priDb "github.com/GoDriveApp/GoDriveApi/domain/contracts/primaryDb"
	"github.com/GoDriveApp/GoDriveApi/domain/errs"
	"github.com/google/uuid"
)

type (
	AccountFactory struct {
		accRepo         priDb.IAccountRepository
		passwordService IPasswordService
	}

	IAccountFactory interface {
		CreateNewAccount(username val.Username, email val.Email, password val.Password) *entities.Account
	}
)

func NewAccountFactory(accRepo priDb.IAccountRepository, passwordService IPasswordService) *AccountFactory {
	return &AccountFactory{accRepo, passwordService}
}

func (acf *AccountFactory) CreateNewAccount(username val.Username, email val.Email, password val.Password) (error, *entities.Account) {
	if acf.accRepo.IsUsernameExist(username) {
		return errs.ThrowDuplicatedUsernameErr(username), nil
	}
	if acf.accRepo.IsEmailExist(email) {
		return errs.ThrowDuplicatedEmailErr(email), nil
	}

	passwordHash := acf.passwordService.Hash(password)
	uid := uuid.New().String()

	return nil, entities.NewAccount(uid, username, email, passwordHash)
}
