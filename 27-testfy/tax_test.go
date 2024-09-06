package tax

import (
	"errors"
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

func TestCalculateAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("Error Save"))
	err := CalculateTaxAndSave(1000, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "Error Save")

	repository.AssertExpectations(t)

}
