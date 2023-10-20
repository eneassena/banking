package service

import (
	"time"

	"github.com/eneassena10/banking/domain"
	"github.com/eneassena10/banking/dto"
	"github.com/eneassena10/banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(*dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.ApiError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.ApiError)
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

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.ApiError) {
	// incoming request validation
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.AccountRepositoryDb.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}
	// if all is well, build the domain object & save the transaction
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.AccountRepositoryDb.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}

func NewAccountService(accountRepositoryDb domain.AccountRepositoryDb) DefaultAccountService {
	return DefaultAccountService{accountRepositoryDb}
}
