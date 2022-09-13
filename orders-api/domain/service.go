package domain

//import "github.com/diego3/kafka-microservices/orders-api/domain/event"

type OrderService struct {
	repository Repository
	//Producer   event.OrderCreatedEventProducer
}

func NewOrderService(repository Repository) *OrderService {
	return &OrderService{
		repository: repository,
	}
}

func (s *OrderService) NewOrder(order Order) {
	//todo business validations ?

	s.repository.PlaceNewOrder(order)
}
