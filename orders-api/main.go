package main

import (
	"log"
	"net/http"

	"github.com/diego3/kafka-microservices/lib"
	"github.com/diego3/kafka-microservices/orders-api/controller"
)

func init() {

}

func main() {
	lib.Hello("Diego")
	http.HandleFunc("/order", controller.HandleCreateOrder)
	err := http.ListenAndServe(":5052", nil)
	if err != nil {
		log.Fatalln("http server error", err)
	}
}
