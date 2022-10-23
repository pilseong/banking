package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pilseong/banking/domain"
	"github.com/pilseong/banking/service"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined")
	}
}

func Start() {

	sanityCheck()

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// ch := CustomerHandlers{
	// 	service: service.NewCustomerService(domain.NewCustomerRepositoryStub()),
	// }
	ch := CustomerHandlers{
		service: service.NewCustomerService(domain.NewCustomerRepositoryDB()),
	}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id}", ch.getCustomer).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	// starting server
	infoString := fmt.Sprintf("%s:%s", address, port)
	log.Printf("The server is running on " + infoString)
	log.Fatal(http.ListenAndServe(infoString, router))
}

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])
// }

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "cool!")
}
