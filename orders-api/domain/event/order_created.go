package event

import "github.com/diego3/kafka-microservices/orders-api/domain"

type OrderCreatedEventProducer interface {
	Produce(order domain.Order)
}
