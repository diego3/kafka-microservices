package event

// https://github.com/segmentio/kafka-go

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	brokers []string
}

func NewKafkaConsumer(brokers []string) *KafkaConsumer {
	return &KafkaConsumer{
		brokers: brokers,
	}
}

func (c *KafkaConsumer) Consume(topic string, partition int) {
	log.Println("starting kafka consumer")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   c.brokers,
		Topic:     topic,
		Partition: partition,

		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	r.SetOffset(0)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("failed to read message: ", err)
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Println("failed to close readeer:", err)
	}
}
