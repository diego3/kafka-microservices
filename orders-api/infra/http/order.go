package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/diego3/kafka-microservices/orders-api/application/command"
	"github.com/diego3/kafka-microservices/orders-api/infra/db"
	"github.com/labstack/echo/v4"
)

func HandleCreateOrder(c echo.Context) error {
	order := command.NewOrderRequest{}
	if err := c.Bind(order); err != nil {
		log.Println("Error trying to bind NewOrderRequest", err)
		return err
	}

	err := command.NewPlaceOrderCommand().Handle(order)
	if err != nil {
		log.Println("Error trying to place new order", err)
		return err
	}
	return c.String(http.StatusOK, "new order placed successfully")
}

func HandleGetOrders(c echo.Context) error {
	repo := db.MongoDB{
		Uri:     "mongodb://localhost:27017",
		Timeout: 3 * time.Second,
	}
	repo.Connection()

	// TODO: serialize a list of orders
	return c.String(http.StatusOK, "[{}]")
}
