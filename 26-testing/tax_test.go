package tax

import "testing"

func TestTax(t *testing.T) {
	amount := 500.0
	expected := 6.0

	result := CalculateTax(amount)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}

}
