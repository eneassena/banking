package domain

import "github.com/eneassena10/banking/errs"

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.ApiError)
	ById(customer_id string) (*Customer, *errs.ApiError)
}
