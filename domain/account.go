package domain

import (
	"github.com/eneassena10/banking/dto"
	"github.com/eneassena10/banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponse() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		CustomerId: a.AccountId,
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount < amount
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.ApiError)
	SaveTransaction(t Transaction) (*Transaction, *errs.ApiError)
	FindBy(accountId string) (*Account, *errs.ApiError)
}
