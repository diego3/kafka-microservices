package service

import "github.com/diego3/kafka-microservices/orders-api/repository"

type Order struct {
	Repository repository.Repository
}
