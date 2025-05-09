package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupEnv() {
	os.Setenv("EXAMPLE_EMPTY_STRING", "")
	os.Setenv("EXAMPLE_STRING", "golang")
	os.Setenv("EXAMPLE_BOOL_FALSE", "0")
	os.Setenv("EXAMPLE_BOOL_TRUE", "1")
	os.Setenv("EXAMPLE_INT", "64")
	os.Setenv("EXAMPLE_SLICE_OF_MAP_OF_STRING", `{"foo":"bar"}`)
	os.Setenv("EXAMPLE_SLICE_OF_JSON_SLICES_OF_STRING", `["foo","bar"]`)
	os.Setenv("EXAMPLE_SLICE_OF_JSON_SLICES_OF_ANY", `["foo","bar",1]`)
	os.Setenv("EXAMPLE_SLICE_OF_JSON_MAP_OF_ANY", `{"foo":"bar","status":true,"value":1}`)
}

func TestGetEnvValue(t *testing.T) {
	setupEnv()
	assert.Equal(t, "ABC", GetEnvValue("EXAMPLE_EMPTY_STRING", "string", "ABC"), "they should be equal")
	assert.Equal(t, "test", GetEnvValue("TEST_ENV", "string", "test"), "they should be equal")
	assert.Equal(t, "golang", GetEnvValue("EXAMPLE_STRING", "string", "test"), "they should be equal")
	assert.Equal(t, false, GetEnvValue("EXAMPLE_BOOL_FALSE", "bool", false), "they should be equal")
	assert.Equal(t, true, GetEnvValue("EXAMPLE_BOOL_TRUE", "bool", nil), "they should be equal")
	assert.Equal(t, 64, GetEnvValue("EXAMPLE_INT", "int", 25), "they should be equal")
	assert.Equal(t, int32(64), GetEnvValue("EXAMPLE_INT", "int32", 25), "they should be equal")
	assert.Equal(t, int64(64), GetEnvValue("EXAMPLE_INT", "int64", 25), "they should be equal")
	assert.Equal(t, float32(64), GetEnvValue("EXAMPLE_INT", "float32", 25), "they should be equal")
	assert.Equal(t, float64(64), GetEnvValue("EXAMPLE_INT", "float64", 25), "they should be equal")
	assert.Equal(t, map[string]string{"foo": "bar"}, GetEnvValue("EXAMPLE_SLICE_OF_MAP_OF_STRING", "string_map_string", nil), "they should be equal")
	assert.Equal(t, []interface{}{"foo", "bar"}, GetEnvValue("EXAMPLE_SLICE_OF_JSON_SLICES_OF_STRING", "json_slice", nil), "they should be equal")
	assert.Equal(t, []interface{}{"foo", "bar", float64(1)}, GetEnvValue("EXAMPLE_SLICE_OF_JSON_SLICES_OF_ANY", "json_slice", nil), "they should be equal")
	assert.Equal(t, map[string]interface{}{"foo": "bar", "status": true, "value": float64(1)}, GetEnvValue("EXAMPLE_SLICE_OF_JSON_MAP_OF_ANY", "json_map_any", nil), "they should be equal")
}
