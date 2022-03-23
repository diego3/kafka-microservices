package main

import (
	"log"
	"net/http"

	"github.com/diego3/kafka-microservices/lib"
	"github.com/diego3/kafka-microservices/orders-api/controller"
	"github.com/diego3/kafka-microservices/orders-api/event"
)

func init() {

}

func main() {
	lib.Hello("Diego")
	log.Println("order-api v1")

	http.HandleFunc("/order", controller.HandleCreateOrder)

	// TODO: blocking call avoided with the goroutine, check if its is a good practice
	// have a consumer and producer at same application
	go event.NewKafkaConsumer().Consume([]string{"localhost:9092"}, "topic-A", 0)

	// never reaching here without using goroutine at kakfa consumer!!!!
	err := http.ListenAndServe(":5052", nil)
	if err != nil {
		log.Fatalln("http server error", err)
	}
}
