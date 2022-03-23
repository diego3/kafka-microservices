package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/diego3/kafka-microservices/orders-api/event"
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
