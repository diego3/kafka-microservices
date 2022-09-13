package db

import "github.com/diego3/kafka-microservices/orders-api/domain"

type MongoOrderRepository struct {
	Db MongoDB
}

func (r *MongoOrderRepository) PlaceNewOrder(order domain.Order) {

}

func (r *MongoOrderRepository) FindById(id string) (domain.Order, error) {
	order := OrderModel{}
	theOrder := fromDocumentModelToOrderDomain(order)
	return *theOrder, nil
}
