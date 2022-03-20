package domain

import "testing"

//go test domain/* -v

func TestCartTotalAmountCalculation(t *testing.T) {
	cart := Cart{}
	phone := Item{Description: "Phone1", Quantity: 1, UnitValue: 1200.0}
	cart.AddItem(phone)

	result := cart.CalcTotalAmount()
	expected := 1200.0
	if result != expected {
		t.Errorf("Result differs from expected %f", expected)
	}
}
