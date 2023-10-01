package service

import (
	"github.com/eneassena10/banking/domain"
	"github.com/eneassena10/banking/domain/dto"
	"github.com/eneassena10/banking/errs"
)

type AccountService interface {
	NewAccount(*dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.ApiError)
}

type DefaultAccountService struct {
	AccountRepositoryDb domain.AccountRepositoryDb
}

func (d DefaultAccountService) NewAccount(account *dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.ApiError) {
	return &dto.NewAccountResponse{}, nil
}

func NewAccountService(accountRepositoryDb domain.AccountRepositoryDb) DefaultAccountService {
	return DefaultAccountService{
		AccountRepositoryDb: accountRepositoryDb,
	}
}
