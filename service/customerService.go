package service

import (
	"github.com/eneassena10/banking/domain"
	"github.com/eneassena10/banking/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.ApiError)
	GetCustomer(string) (*domain.Customer, *errs.ApiError)
	GetByStatusCustomer(customer_status string) ([]domain.Customer, *errs.ApiError)
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

func (s DefaultCustomerService) GetByStatusCustomer(customer_status string) ([]domain.Customer, *errs.ApiError) {
	return s.repo.FindByStatus(customer_status)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
