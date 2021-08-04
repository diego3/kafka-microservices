package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"go.customer/model"
)

var customers []model.Customer

type CustomerService struct {
}

func New() *CustomerService {
	return &CustomerService{}
}

func (c *CustomerService) GetAll() []model.Customer {
	return customers
}

func (c *CustomerService) Add(customer model.Customer) (model.Customer, error) {
	customer.Id = uuid.NewString()
	validation := c.IsValid(customer)
	if !validation.Status {
		return model.Customer{}, errors.New("invalid customer " + validation.Error)
	}
	customers = append(customers, customer)
	return customer, nil
}

func (c *CustomerService) GetById(customerId string) model.Customer {
	if len(customers) == 0 {
		return model.Customer{}
	}

	for _, r := range customers {
		if strings.Compare(r.Id, customerId) == 0 {
			return r
		}
	}
	return model.Customer{}
}

func (c *CustomerService) GetByEmail(email string) model.Customer {
	if len(customers) == 0 {
		return model.Customer{}
	}

	for _, r := range customers {
		if strings.Compare(r.Email, email) == 0 {
			return r
		}
	}
	return model.Customer{}
}

func (c *CustomerService) IsValid(in model.Customer) model.Validation {
	customer := c.GetByEmail(in.Email)
	if strings.Compare(customer.Id, "") != 0 {
		return model.Validation{Status: false, Error: "Email already exists"}
	}
	return model.Validation{Status: true}
}
