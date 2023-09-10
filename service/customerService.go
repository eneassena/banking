package service

import (
	"github.com/eneassena10/banking/domain"
	"github.com/eneassena10/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *errs.ApiError)
	GetCustomer(customer_id string) (*domain.Customer, *errs.ApiError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.ApiError) {
	return s.repo.FindAll(StatusMap[status])
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.ApiError) {
	return s.repo.ById(id)
}
