package domain

type OrderStatus int64

const (
	PENDING OrderStatus = iota
	APPROVED
	REJECTED
)

func (s OrderStatus) ToString() string {
	switch s {
	case 0:
		return "PENDING"
	case 1:
		return "APPROVED"
	case 2:
		return "REJECTED"
	default:
		return "N/A"
	}
}

type Order struct {
	id     string
	email  string
	amount float64
	status OrderStatus
}

func NewOrder(id, email string, amount float64) *Order {
	return &Order{
		id:     id,
		email:  email,
		amount: amount,
		status: PENDING,
	}
}

func (o *Order) Approve() error {
	// validate when order can be approved or not
	return nil
}

func (o *Order) GetId() string {
	return o.id
}

func (o *Order) GetEmail() string {
	return o.email
}

func (o *Order) GetAmount() float64 {
	return o.amount
}

func (o *Order) GetStatus() OrderStatus {
	return o.status
}
