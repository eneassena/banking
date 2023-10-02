package service

import (
	"time"

	"github.com/eneassena10/banking/domain"
	"github.com/eneassena10/banking/dto"
	"github.com/eneassena10/banking/errs"
)

type AccountService interface {
	NewAccount(*dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.ApiError)
}

type DefaultAccountService struct {
	AccountRepositoryDb domain.AccountRepositoryDb
}

func (d DefaultAccountService) NewAccount(req *dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.ApiError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		CustomerId:  "",
		AccountId:   req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	newAccount, err := d.AccountRepositoryDb.Save(a)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponse()

	return &response, nil
}

func NewAccountService(accountRepositoryDb domain.AccountRepositoryDb) DefaultAccountService {
	return DefaultAccountService{accountRepositoryDb}
}
