package variable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getVariableWithInterfaceType(kind string) interface{} {
	var ret interface{}
	switch kind {
	case "string":
		ret = "hello"
	case "bool":
		ret = true
	case "int":
		ret = 10
	case "int32":
		ret = int32(10)
	case "int64":
		ret = int64(10)
	}
	return ret
}
func TestCheckType(t *testing.T) {
	assert.Equal(t, "string", CheckType("hello"))
	assert.Equal(t, "bool", CheckType(true))
	assert.Equal(t, "int", CheckType(10))
	assert.Equal(t, "int32", CheckType(int32(10)))
	assert.Equal(t, "int64", CheckType(int64(10)))
	assert.Equal(t, "float32", CheckType(float32(10)))
	assert.Equal(t, "float64", CheckType(float64(10)))
	assert.Equal(t, "[]string", CheckType([]string{"hello", "world"}))
	assert.Equal(t, "[]bool", CheckType([]bool{true, false}))
	assert.Equal(t, "[]int", CheckType([]int{10, 20}))
	assert.Equal(t, "[]int32", CheckType([]int32{10, 20}))
	assert.Equal(t, "[]int64", CheckType([]int64{10, 20}))
	assert.Equal(t, "[]float32", CheckType([]float32{10, 20}))
	assert.Equal(t, "[]float64", CheckType([]float64{10, 20}))
	assert.Equal(t, "[]float64", CheckType([]float64{10, 20}))
	assert.Equal(t, "[]interface{}", CheckType([]interface{}{getVariableWithInterfaceType("string"), getVariableWithInterfaceType("bool")}))
	assert.Equal(t, "unknown", CheckType(map[string]string{"hello": "world"}))
}
