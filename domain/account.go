package domain

import (
	"github.com/eneassena10/banking/dto"
	"github.com/eneassena10/banking/errs"
)

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

type AccountRepository interface {
	Save(Account) (*Account, *errs.ApiError)
}
