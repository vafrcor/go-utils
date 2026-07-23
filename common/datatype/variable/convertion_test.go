package variable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertSoifIntoSos(t *testing.T) {
	a := []interface{}{"foo", "bar"}
	assert.Equal(t, []string{"foo", "bar"}, ConvertSoifIntoSos(a))
}
