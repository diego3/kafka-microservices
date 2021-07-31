package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.customer/service"
)

func AliveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"alive\": \"true\"}")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	//id := uuid.NewString()
	//c1 := model.Customer{Id: id, Name: "Diego", Email: "diego@gmail.com"}
	customerService := service.CustomerService{}
	customers := customerService.GetAll()
	json.NewEncoder(w).Encode(customers)
}
