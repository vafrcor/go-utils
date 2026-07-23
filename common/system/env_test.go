package env

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupEnvForTestGetEnvValue() {
	os.Setenv("EXAMPLE_EMPTY_STRING", "")
	os.Setenv("EXAMPLE_STRING", "golang")
	os.Setenv("EXAMPLE_BOOL_FALSE", "0")
	os.Setenv("EXAMPLE_BOOL_TRUE", "1")
	os.Setenv("EXAMPLE_INT", "64")
	os.Setenv("EXAMPLE_SLICE_OF_MAP_OF_STRING", `{"foo":"bar"}`)
	os.Setenv("EXAMPLE_SLICE_OF_JSON_SLICES_OF_STRING", `["foo","bar"]`)
	os.Setenv("EXAMPLE_SLICE_OF_JSON_SLICES_OF_ANY", `["foo","bar",1]`)
	os.Setenv("EXAMPLE_SLICE_OF_JSON_MAP_OF_ANY", `{"foo":"bar","status":true,"value":1}`)
	os.Setenv("EXAMPLE_JSON_MAP_OF_ANY", `{"foo":"bar","status":true,"value":1}`)
	os.Setenv("EXAMPLE_JSON_MAP_OF_ANY_2", `{"abc":15,"def":10,"ghi":5,"jkl":3}`)
}

func TestGetEnvValue(t *testing.T) {
	setupEnvForTestGetEnvValue()
	testCases := []struct {
		Env      string
		Type     string
		Expected interface{}
		Default  interface{}
		Message  string
	}{
		{"EXAMPLE_EMPTY_STRING", "string", "ABC", "ABC", "they should be equal"},
		{"TEST_ENV", "string", "test", "test", "they should be equal"},
		{"EXAMPLE_STRING", "string", "golang", "test", "they should be equal"},
		{"EXAMPLE_BOOL_FALSE", "bool", false, false, "they should be equal"},
		{"EXAMPLE_BOOL_TRUE", "bool", true, nil, "they should be equal"},
		{"EXAMPLE_INT", "int", 64, 25, "they should be equal"},
		{"EXAMPLE_INT", "int32", int32(64), 25, "they should be equal"},
		{"EXAMPLE_INT", "int64", int64(64), 25, "they should be equal"},
		{"EXAMPLE_INT", "uint64", uint64(64), 25, "they should be equal"},
		{"EXAMPLE_INT", "float32", float32(64), 25, "they should be equal"},
		{"EXAMPLE_INT", "float64", float64(64), 25, "they should be equal"},
		{"EXAMPLE_SLICE_OF_MAP_OF_STRING", "string_map_string", map[string]string{"foo": "bar"}, nil, "they should be equal"},
		{"EXAMPLE_SLICE_OF_JSON_SLICES_OF_STRING", "json_slice", []interface{}{"foo", "bar"}, nil, "they should be equal"},
		{"EXAMPLE_SLICE_OF_JSON_SLICES_OF_ANY", "json_slice", []interface{}{"foo", "bar", float64(1)}, nil, "they should be equal"},
		{"EXAMPLE_JSON_MAP_OF_ANY", "json_map_any", map[string]interface{}{"foo": "bar", "status": true, "value": float64(1)}, nil, "they should be equal"},
		{"EXAMPLE_JSON_MAP_OF_ANY_2", "json_map_any", map[string]interface{}{"abc": float64(15), "def": float64(10), "ghi": float64(5), "jkl": float64(3)}, nil, "they should be equal"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("expect:%s", tc.Env), func(t *testing.T) {
			got := GetEnvValue(tc.Env, tc.Type, tc.Default)
			assert.Equal(t, tc.Expected, got, tc.Message)
		})
	}
}

func TestConvertEnvDurationStringToTimeDuration(t *testing.T) {
	testCases := []struct {
		Env      string
		Expected time.Duration
		Default  string
	}{
		{"TEST_DURATION_1_MINUTE", 1 * time.Minute, "1m"},
		{"TEST_DURATION_1_SECOND", 1 * time.Second, "1s"},
		{"TEST_DURATION_N5_SECOND", -5 * time.Second, "-5s"},
		{"TEST_DURATION_100_MILLISECOND", 100 * time.Millisecond, "100ms"},
		{"TEST_DURATION_100_MICROSECOND", 100 * time.Microsecond, "100us"},
		{"TEST_DURATION_100_NANOSECOND", 100 * time.Nanosecond, "100ns"},
		{"TEST_DURATION_1_HOUR", 1 * time.Hour, "1h"},
		{"TEST_DURATION_2_HOUR", 2 * time.Hour, "2h"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("expect:%s", tc.Env), func(t *testing.T) {
			got := ConvertEnvDurationStringToTimeDuration(tc.Env, tc.Default)
			assert.Equal(t, tc.Expected, got)
		})
	}
}
