package lib

import (
	"context"
	"fmt"
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
	r.SetOffset(kafka.LastOffset)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("failed to read message: ", err)
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Println("failed to close reader:", err)
	}
}
