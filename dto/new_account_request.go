package dto

import (
	"strings"

	"github.com/eneassena10/banking/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r *NewAccountRequest) Validate() *errs.ApiError {
	const (
		Saving   = "saving"
		Checking = "checking"
	)
	if r.Amount > 5000 {
		return errs.NewValidationError("To open a new account you need to deposit atleast 5000.00")
	}
	if strings.ToLower(r.AccountType) != Saving && strings.ToLower(r.AccountType) != Checking {
		return errs.NewValidationError("Account type should be checking or serving")
	}
	return nil
}
