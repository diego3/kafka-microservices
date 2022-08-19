package event

// https://github.com/segmentio/kafka-go

type OrderConsumer interface {
	Consume()
}
