package app

import (
	"log"
	"net/http"

	"github.com/eneassena10/banking/domain"
	"github.com/eneassena10/banking/errs"
	"github.com/eneassena10/banking/service"
	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

var StatusMap = map[string]string{
	"active":   "1",
	"inactive": "0",
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	queryURL := r.URL.Query().Get("status")
	log.Println("queryURL: " + StatusMap[queryURL])

	var customers []domain.Customer
	var err *errs.ApiError

	customer_status := StatusMap[queryURL]
	if queryURL != "" {
		customers, err = ch.service.GetByStatusCustomer(customer_status)
		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
			return
		}
		writeResponse(w, http.StatusOK, customers)
		return
	}

	customers, err = ch.service.GetAllCustomer()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(customer_id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}
