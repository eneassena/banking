package domain

import (
	"database/sql"
	"time"

	"github.com/eneassena10/banking/errs"
	"github.com/eneassena10/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:Jose123_+#@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		logger.Error(err.Error())
	}

	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	if err := client.Ping(); err != nil {
		logger.Error(err.Error())
	}
	return CustomerRepositoryDb{client: client}
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
