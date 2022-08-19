package kafka

import "github.com/diego3/kafka-microservices/lib"

type OrderConsumer struct {
}

func (c *OrderConsumer) Consume() {
	var brokers []string
	var partition int
	lib.NewKafkaConsumer(brokers).Consume("topic", partition)
}
