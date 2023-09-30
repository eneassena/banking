package service

import (
	"github.com/eneassena10/banking/domain"
	"github.com/eneassena10/banking/domain/dto"
	"github.com/eneassena10/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *errs.ApiError)
	GetCustomer(customer_id string) (*dto.CustomerResponse, *errs.ApiError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.ApiError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	customerResponse := c.ToDto()

	return customerResponse, nil
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.ApiError) {
	return s.repo.FindAll(StatusMap[status])
}
