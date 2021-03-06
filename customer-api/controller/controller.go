package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/diego3/kafka-microservices/customer-api/model"
	"github.com/diego3/kafka-microservices/customer-api/service"
)

func AliveHandler(w http.ResponseWriter, r *http.Request) {
	//ctx := apm.DefaultTracer.StartTransaction("GET /alive", "request")
	//defer ctx.End()ELASTIC_APM_SERVER_URL
	fmt.Fprintf(w, "{\"alive\": \"true\"}")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := service.New().GetAll()
	json.NewEncoder(w).Encode(customers)
}

func RegisterNewCustomer(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)
	var customer = model.Customer{}
	json.Unmarshal(bytes, &customer)

	newCustomer, err := service.New().Add(customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := model.ErrorResponse{Message: err.Error()}
		jsonBytes, _ := json.Marshal(response)
		fmt.Fprintf(w, string(jsonBytes))
		log.Println("error trying to register new customer")
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newCustomer)
		log.Println("new customer OK")
	}

	log.Println("ELASTIC_APM_SERVER_URL=" + os.Getenv("ELASTIC_APM_SERVER_URL"))
}
