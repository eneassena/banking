package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func writeResponseWithXml(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(code)
	if err := xml.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
