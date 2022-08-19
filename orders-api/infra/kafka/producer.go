package kafka

import (
	"encoding/json"
	"log"

	"github.com/diego3/kafka-microservices/lib"
	"github.com/diego3/kafka-microservices/orders-api/domain"
)

type KafkaOrderCreatedEventProducer struct {
}

func (k *KafkaOrderCreatedEventProducer) Produce(order domain.Order) {
	bytes, err := json.Marshal(order)
	if err != nil {
		log.Println("error trying to marshal some order", err)
		return
	}

	lib.NewKafkaProducer().Produce("localhost:9092", "ecommerce-new-order", "order", string(bytes))
}
