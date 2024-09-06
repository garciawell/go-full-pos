package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	tax, err := CalculateTax(1000.00)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0)
	assert.Error(t, err, "Amount must be grather than 0")
	assert.Equal(t, 0.0, tax)
}
