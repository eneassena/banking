package domain

import (
	"database/sql"

	"github.com/eneassena10/banking/errs"
	"github.com/eneassena10/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.ApiError) {
	customers := make([]Customer, 0)
	var err error

	if status == "" {
		findAllQuery := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customers"
		err = d.client.Select(&customers, findAllQuery)
	} else {
		findAllQuery := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customers where status=?"
		err = d.client.Select(&customers, findAllQuery, status)
	}

	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.ApiError) {
	customerSql := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customers where customer_id=?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Customer not found, " + err.Error())
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer, " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client: dbClient}
}
