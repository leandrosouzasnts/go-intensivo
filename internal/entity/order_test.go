package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGesAnErrorIfIdIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "id is required")
}

func TestFinalPrice(t *testing.T) {
	order := Order{Id: "1", Price: 10.0, Tax: 1.0}

	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
