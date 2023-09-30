package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/eneassena10/banking/domain"
	"github.com/eneassena10/banking/logger"
	"github.com/eneassena10/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable is not defined")
	}
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	// wiring
	dbClient := getDbConnect()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}

	// accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	// define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info("Starting the application http://127.0.0.1:8000")
	http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)
}

func getDbConnect() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		logger.Error(err.Error())
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	if err = client.Ping(); err != nil {
		logger.Error(err.Error())
	}

	return client
}
