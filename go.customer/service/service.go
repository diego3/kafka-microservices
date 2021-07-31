package service

import (
	"github.com/google/uuid"
	"go.customer/model"
)

type CustomerService struct {
}

func (c *CustomerService) GetAll() []model.Customer {
	customers := []model.Customer{
		{Id: uuid.NewString(), Name: "Jose", Email: "email1@gmail.com"},
		{Id: uuid.NewString(), Name: "Paulo", Email: "email2@gmail.com"},
		{Id: uuid.NewString(), Name: "Pedro", Email: "email3@gmail.com"},
	}
	return customers
}
