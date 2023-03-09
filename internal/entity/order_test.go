package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorIfIdIsBlank(testing *testing.T) {
	order := Order{}

	assert.Error(testing, order.Validate(), "invalid id")
}

func TestIfItGetsAnErrorIfPriceIsBlank(testing *testing.T) {
	order := Order{Id: "123"}

	assert.Error(testing, order.Validate(), "invalid price")
}

func TestIfItGetsAnErrorIfTaxIsBlank(testing *testing.T) {
	order := Order{Id: "123", Price: 10.0}

	assert.Error(testing, order.Validate(), "invalid tax")
}

func TestWithAllValidParams(testing *testing.T) {
	order := Order{Id: "123", Price: 10.0, Tax: 1.0}
	assert.NoError(testing, order.Validate())
	assert.Equal(testing, 10.0, order.Price)
	assert.Equal(testing, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(testing, 11.0, order.FinalPrice)
	order.CalculateFinalPriceWithDiscount(10.0)
	assert.Equal(testing, 1.1, order.Discount)
}
