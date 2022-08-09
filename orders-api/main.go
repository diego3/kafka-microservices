package main

import (
	"log"
	"net/http"
	"time"

	"github.com/diego3/kafka-microservices/lib"
	controller "github.com/diego3/kafka-microservices/orders-api/infra/http"

	//"github.com/diego3/kafka-microservices/orders-api/event"
	"github.com/gorilla/mux"
)

// go get google.golang.org/grpc
// go get google.golang.org/protobuf
// go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

func init() {

}

func main() {
	lib.Hello("Diego")
	log.Println("order-api v1")

	router := mux.NewRouter()
	router.HandleFunc("/order", controller.HandleCreateOrder).Methods(http.MethodPost)
	router.HandleFunc("/order", controller.HandleGetOrders).Methods(http.MethodGet)

	// TODO: blocking call avoided with the goroutine, check if its is a good practice
	// have a consumer and producer at same application
	//go event.NewKafkaConsumer([]string{"localhost:9092"}).Consume("ecommerce-new-order", 0)

	server := http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:5052",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}
	err := server.ListenAndServe()

	// never reaching here without using goroutine at kakfa consumer!!!!
	if err != nil {
		log.Fatalln("http server error", err)
	}
}
