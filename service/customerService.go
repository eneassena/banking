package service

import (
	"github.com/eneassena10/banking/domain"
	"github.com/eneassena10/banking/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.ApiError)
	GetCustomer(string) (*domain.Customer, *errs.ApiError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.ApiError) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.ApiError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
