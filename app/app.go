package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eneassena10/banking/domain"
	"github.com/eneassena10/banking/logger"
	"github.com/eneassena10/banking/service"
	"github.com/gorilla/mux"
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
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info("Starting the application http://127.0.0.1:8000")
	http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)
}
