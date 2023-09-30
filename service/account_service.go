package service

import (
	"github.com/eneassena10/banking/domain/dto"
	"github.com/eneassena10/banking/errs"
)

type AccountService interface {
	NewAccount(*dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.ApiError)
}

type DefaultAccountService struct{}
