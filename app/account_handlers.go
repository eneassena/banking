package app

import (
	"encoding/json"
	"net/http"

	"github.com/eneassena10/banking/dto"
	"github.com/eneassena10/banking/service"
	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	service service.DefaultAccountService
}

func (a AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_id := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else {
		request.CustomerId = customer_id
		account, appError := a.service.NewAccount(&request)
		if appError != nil {
			writeResponse(w, http.StatusInternalServerError, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

// customer/2000/accounts/21231
