package domain

import (
	"fmt"
	"os"
	"time"

	"github.com/eneassena10/banking/logger"
	"github.com/jmoiron/sqlx"
)

type ConnectMysql struct {
	client *sqlx.DB
}

func (d ConnectMysql) DBClient() *sqlx.DB {
	if d.client == nil {
		var (
			err        error
			dbUser     string = os.Getenv("DB_USER")
			dbPassword string = os.Getenv("DB_PASSWORD")
			// dbPassword   string = "root"
			dbAddr string = os.Getenv("DB_ADDR")
			dbPort string = os.Getenv("DB_PORT")
			dbName string = os.Getenv("DB_NAME")
		)

		dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)
		d.client, err = sqlx.Open("mysql", dataSource)
		if err != nil {
			logger.Error(err.Error())
		}

		d.client.SetConnMaxLifetime(time.Minute * 3)
		d.client.SetMaxOpenConns(10)
		d.client.SetMaxIdleConns(10)
		if err = d.client.Ping(); err != nil {
			logger.Error(err.Error())
		}
	}
	return d.client
}
