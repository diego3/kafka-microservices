package domain

import (
	"fmt"
	"testing"
)

//go test domain/* -v

func TestCartTotalAmountCalculation(t *testing.T) {
	assertExcepted := func(t *testing.T, result, expected float64) {
		t.Helper()
		if result != expected {
			t.Errorf("Result: %f  from expected %f", result, expected)
		}
	}

	t.Run("Given 0 item should calculate zero", func(t *testing.T) {
		cart := Cart{}
		result := cart.CalcTotalAmount()
		expected := 0.0
		assertExcepted(t, result, expected)
	})

	t.Run("Given 1 item should calculate successfully", func(t *testing.T) {
		cart := Cart{}
		phone := Item{Description: "Phone1", Quantity: 1, UnitValue: 1200.0}
		cart.AddItem(phone)

		result := cart.CalcTotalAmount()
		expected := 1200.0
		assertExcepted(t, result, expected)
	})

	t.Run("Given 2 items should calculate successfully", func(t *testing.T) {
		cart := Cart{}
		phone := Item{Description: "Phone1", Quantity: 1, UnitValue: 1200.0}
		smartTv := Item{Description: "SmartTv", Quantity: 1, UnitValue: 4000.0}
		cart.AddItem(phone)
		cart.AddItem(smartTv)

		result := cart.CalcTotalAmount()
		expected := 5200.0
		assertExcepted(t, result, expected)
	})

	t.Run("Given items with negative values should calculate successfully", func(t *testing.T) {
		cart := Cart{}
		phone := Item{Description: "Phone1", Quantity: 1, UnitValue: -1000.0}
		smartTv := Item{Description: "SmartTv", Quantity: 1, UnitValue: 1000.0}
		cart.AddItem(phone)
		cart.AddItem(smartTv)

		result := cart.CalcTotalAmount()
		expected := 2000.0
		assertExcepted(t, result, expected)
	})

	t.Run("Given items with negative quantities should calculate successfully", func(t *testing.T) {
		cart := Cart{}
		phone := Item{Description: "Phone1", Quantity: 1, UnitValue: -1000.0}
		smartTv := Item{Description: "SmartTv", Quantity: -3, UnitValue: 1000.0}
		cart.AddItem(phone)
		cart.AddItem(smartTv)

		result := cart.CalcTotalAmount()
		expected := 4000.0
		assertExcepted(t, result, expected)
	})

	t.Run("Given items with different quantities", func(t *testing.T) {
		cart := Cart{}
		phone := Item{Description: "Phone1", Quantity: 2, UnitValue: 500.0}
		smartTv := Item{Description: "SmartTv", Quantity: 5, UnitValue: 1000.0}
		cart.AddItem(phone)
		cart.AddItem(smartTv)

		result := cart.CalcTotalAmount()
		expected := 6000.0
		assertExcepted(t, result, expected)
	})

}

func ExampleCart_CalcTotalAmount() {
	cart := Cart{}
	phone := Item{Description: "Phone1", Quantity: 2, UnitValue: 500.0}
	smartTv := Item{Description: "SmartTv", Quantity: 5, UnitValue: 1000.0}
	cart.AddItem(phone)
	cart.AddItem(smartTv)

	fmt.Println(cart.CalcTotalAmount())
	// Output: 6000
}
