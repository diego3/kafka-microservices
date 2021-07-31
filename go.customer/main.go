package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.customer/controller"
)

const PORT = ":9090"

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controller.AliveHandler).Methods(http.MethodGet)
	router.HandleFunc("/customer", controller.GetAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(PORT, router))
}

func init() {
	log.Println("Customer service init")
}

func main() {
	log.Println("API Customer v1.0.0")
	handleRequests()
}
