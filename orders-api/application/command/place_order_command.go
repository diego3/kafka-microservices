package command

import (
	"log"
	"strconv"

	"github.com/diego3/kafka-microservices/orders-api/domain"
	"github.com/diego3/kafka-microservices/orders-api/infra/db"
	"github.com/google/uuid"
)

type NewOrderRequest struct {
	email  string `json:"email"`
	amount string `json:"amount"`
}

type PlaceOrderCommand struct {
}

func NewPlaceOrderCommand() *PlaceOrderCommand {
	return &PlaceOrderCommand{}
}

func (c *PlaceOrderCommand) Handle(request NewOrderRequest) error {
	repository := db.MongoOrderRepository{
		Db: db.MongoDB{
			Uri: "localhost:27017",
		},
	}
	service := domain.NewOrderService(&repository)
	amount, err := strconv.ParseFloat(request.amount, 32)
	if err != nil {
		log.Println("Amount parse error", err)
	}
	order := domain.NewOrder(uuid.NewString(), request.email, amount)

	service.NewOrder(*order)

	return nil
}
