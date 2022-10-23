package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pilseong/banking/service"
)

// type Customer struct {
// 	Name    string `json:"full_name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	Zipcode string `json:"zip_code" xml:"zipcode"`
// }

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		writeResponse(w, r, err.Code, err.AsMessage())
	} else {
		// send response following the requested content-type from the client
		writeResponse(w, r, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, r, err.Code, err.AsMessage())
	} else {
		// send response following the requested content-type from the client
		writeResponse(w, r, http.StatusOK, customer)
	}

}

func writeResponse(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application-xml")
		if err := xml.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	} else {
		w.Header().Add("Content-Type", "application-json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			panic(err)
		}
	}
}
