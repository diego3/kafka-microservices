package event

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
}

func NewKafkaProducer() *KafkaProducer {
	return &KafkaProducer{}
}

func (p *KafkaProducer) Produce(address, topic, messageKey, message string) {
	w := kafka.Writer{
		Addr:         kafka.TCP(address),
		Topic:        topic,
		RequiredAcks: kafka.RequireAll,
		Balancer:     &kafka.LeastBytes{},
	}

	msg := kafka.Message{
		Key:   []byte(messageKey),
		Value: []byte(message),
	}

	err := w.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Println("Failed to write message:", err)
	}

	if err := w.Close(); err != nil {
		log.Println("Failed to close writer:", err)
	}
}
