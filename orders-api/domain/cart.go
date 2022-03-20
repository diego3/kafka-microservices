package domain

import "math"

type Item struct {
	Description string
	Quantity    int64
	UnitValue   float64
}

type Cart struct {
	Items []Item
}

// Calculate total amount from a Cart and its items
func (c *Cart) CalcTotalAmount() float64 {
	var total float64 = 0
	if len(c.Items) == 0 {
		return total
	}
	for _, item := range c.Items {
		total += item.CalcTotalAmount()
	}
	return total
}

func (c *Cart) AddItem(item Item) {
	added := append(c.Items, item)
	c.Items = added
}

// Calculate total amount for an Item
func (c *Item) CalcTotalAmount() float64 {
	return math.Abs(c.UnitValue) * math.Abs(float64(c.Quantity))
}
