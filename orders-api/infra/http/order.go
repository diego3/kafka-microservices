package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/diego3/kafka-microservices/orders-api/domain"
	"github.com/diego3/kafka-microservices/orders-api/infra/db"
)

func HandleCreateOrder(resp http.ResponseWriter, req *http.Request) {
	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("CreateOrder Error: trying to read body", err)
		return
	}
	//fmt.Println(string(bytes))
	order := domain.Order{}
	err = json.Unmarshal(bytes, &order)
	if err != nil {
		log.Println("Error trying to marshal payload", err)
		resp.WriteHeader(http.StatusBadRequest)
		io.WriteString(resp, "invalid request")
		return
	}

	repository := db.MongoOrderRepository{
		Db: db.MongoDB{
			Uri: "localhost:27017",
		},
	}
	service := domain.OrderService{
		Repository: &repository,
	}
	service.Repository.PlaceNewOrder(order)

	resp.WriteHeader(http.StatusOK)
	io.WriteString(resp, "ok")
}

func HandleGetOrders(w http.ResponseWriter, req *http.Request) {
	//vars := mux.Vars(r)
	repo := db.MongoDB{
		Uri:     "mongodb://localhost:27017",
		Timeout: 3 * time.Second,
	}
	repo.Connection()

	// TODO: serialize a list of orders
	io.WriteString(w, "[{}]")
}
