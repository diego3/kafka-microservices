package application

import (
	"log"
	"net/http"
	"time"

	controller "github.com/diego3/kafka-microservices/orders-api/infra/http"
	"github.com/labstack/echo/v4"
)

// https://medium.com/avenue-tech/dependency-injection-in-go-35293ef7b6
// https://medium.com/avenue-tech/arquitetura-hexagonal-com-golang-c344411aa940
type WebApp struct {
}

func (w *WebApp) Run() {
	router := echo.New()
	router.POST("/order", controller.HandleCreateOrder)
	router.GET("/order", controller.HandleGetOrders)

	server := http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:5052",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	log.Println("order-api v1")
	err := server.ListenAndServe()

	// never reaching here without using goroutine at kakfa consumer!!!!
	if err != nil {
		log.Fatalln("http server error", err)
	}
}
