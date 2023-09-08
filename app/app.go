package app

import (
	"log"
	"net/http"

	"github.com/eneassena10/banking/domain"
	"github.com/eneassena10/banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
