package domain

type Repository interface {
	PlaceNewOrder(order Order)
	FindById(id string) (Order, error)
}

//func (r *Repository) PlaceNewOrder(order Order) {}
