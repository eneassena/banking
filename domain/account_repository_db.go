package domain

import (
	"strconv"

	"github.com/eneassena10/banking/errs"
	"github.com/eneassena10/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.ApiError) {
	queryInsert := "INSERT INTO accounts (customer_id, opening_date,account_type, amount, status) VALUES (?,?,?,?,?);"
	result, err := d.client.Exec(queryInsert, a.AccountId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating from new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected from error database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected from error database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
