package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundToDecimal(t *testing.T) {
	assert.Equal(t, float64(1.23), RoundToDecimal(1.2345, 2))
	assert.Equal(t, float64(1.00), RoundToDecimal(1.2345, 0))
}

func TestCeilInt(t *testing.T) {
	assert.Equal(t, int(1), CeilInt(1.00))
	assert.Equal(t, int(2), CeilInt(1.23))
	assert.Equal(t, int(2), CeilInt(1.51))
}
