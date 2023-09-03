package domain

import "github.com/eneassena10/banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.ApiError)
	ById(string) (*Customer, *errs.ApiError)
}
