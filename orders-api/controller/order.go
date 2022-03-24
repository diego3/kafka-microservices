package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/diego3/kafka-microservices/orders-api/event"
	"github.com/diego3/kafka-microservices/orders-api/repository"
)

func HandleCreateOrder(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln("CreateOrder Error: trying to read body", err)
	}
	fmt.Println("enviando para fila")
	fmt.Println(string(bytes))

	//TODO: create abstraction here!!!!
	event.NewKafkaProducer().Produce("localhost:9092", "topic-A", "order", string(bytes))

	io.WriteString(w, "ok")
}

func HandleGetOrders(w http.ResponseWriter, req *http.Request) {
	//vars := mux.Vars(r)
	repo := repository.MongoOrderRepository{
		Uri:     "mongodb://localhost:27017",
		Timeout: 3 * time.Second,
	}
	repo.Connection()

	// TODO: serialize a list of orders
	io.WriteString(w, "[{}]")
}
