package tax

import "errors"

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("Amount must be grather than 0")
	}
	if amount >= 1000 && amount < 20000 {
		return 10, nil
	}
	if amount >= 20000 {
		return 20.0, nil
	}
	return 5, nil
}
