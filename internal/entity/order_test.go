package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testIfItGetsAnErrorIfIdIsBlank(testing *testing.T) {
	order := Order{}

	assert.Error(testing, order.validate(), "invalid id")
}

func testIfItGetsAnErrorIfPriceIsBlank(testing *testing.T) {
	order := Order{Id: "123"}

	assert.Error(testing, order.validate(), "invalid price")
}

func testIfItGetsAnErrorIfTaxIsBlank(testing *testing.T) {
	order := Order{Id: "123", Price: 10.0}

	assert.Error(testing, order.validate(), "invalid tax")
}

func testWithAllValidParams(testing *testing.T) {
	order := Order{Id: "123", Price: 10.0, Tax: 1.0}
	assert.Error(testing, order.validate())
	assert.Equal(testing, 10.0, order.Price)
	assert.Equal(testing, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(testing, 11.0, order.FinalPrice)
}
