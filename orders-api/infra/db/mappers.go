package db

import "github.com/diego3/kafka-microservices/orders-api/domain"

func fromOrderDomainToDocumentModel(order domain.Order) *OrderModel {
	model := OrderModel{
		ID:     order.GetId(),
		Email:  order.GetEmail(),
		Amount: order.GetAmount(),
		Status: order.GetStatus().ToString(),
	}
	return &model
}

func fromDocumentModelToOrderDomain(model OrderModel) *domain.Order {
	// qual o jeito ideal de carregar o status sem usar um setter anêmico?
	return domain.NewOrder(model.ID, model.Email, model.Amount)
}
