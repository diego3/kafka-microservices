package main

import (
	//"context"
	"log"
	"net/http"

	//"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"go.customer/controller"
	"go.elastic.co/apm/module/apmgorilla"
)

const PORT = ":9090"

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controller.AliveHandler).Methods(http.MethodGet)
	router.HandleFunc("/customer", controller.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer", controller.RegisterNewCustomer).Methods(http.MethodPost)

	apmgorilla.Instrument(router)

	log.Fatal(http.ListenAndServe(PORT, router)) //apmhttp.Wrap(router)
}

func init() {
	log.Println("Customer service init")
	/*redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	redisClient.LPush(context.Background(), "redis-key-1", "customer1")
	log.Println(redisClient)*/
}

func main() {
	log.Println("API Customer v1.0.0")
	handleRequests()
}
