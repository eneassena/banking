package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/eneassena10/banking/errs"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:Jose123_+#@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	if err := client.Ping(); err != nil {
		log.Fatal(err.Error())
	}
	return CustomerRepositoryDb{client: client}
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.ApiError) {
	findAllQuery := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customers"

	rows, err := d.client.Query(findAllQuery)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		customer := Customer{}
		if err := rows.Scan(&customer.Id, &customer.Name, &customer.DateOfBirth, &customer.City, &customer.ZipCode, &customer.Status); err != nil {
			return nil, errs.NewUnexpectedError(err.Error())
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.ApiError) {
	customerSql := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customers where customer_id=?"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDb) FindByStatus(customer_status string) ([]Customer, *errs.ApiError) {
	customerSql := "SELECT customer_id,name,date_of_birth,city,zipcode,status FROM customers where status=?"

	rows, err := d.client.Query(customerSql, customer_status)
	if err != nil {
		return nil, errs.NewUnexpectedError("Error database Unexpected")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		customer := Customer{}
		if err := rows.Scan(&customer.Id, &customer.Name, &customer.DateOfBirth, &customer.City, &customer.ZipCode, &customer.Status); err != nil {
			return nil, errs.NewUnexpectedError("Error scanning customer")
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
