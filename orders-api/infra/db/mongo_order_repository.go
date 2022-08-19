package db

import "github.com/diego3/kafka-microservices/orders-api/domain"

type MongoOrderRepository struct {
	Db MongoDB
}

func (r *MongoOrderRepository) PlaceNewOrder(order domain.Order) {
	// map domain to document
}

func (r *MongoOrderRepository) FindById(id string) (domain.Order, error) {
	order := domain.Order{}
	// map document to domain
	return order, nil
}
