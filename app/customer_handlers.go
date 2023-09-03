package app

import (
	"net/http"

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

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		writeResponseWithXml(w, http.StatusOK, customers)
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
