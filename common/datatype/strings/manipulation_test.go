package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringLeftPad(t *testing.T) {
	assert.Equal(t, "01", StringLeftPad("1", 2, "0"))
	assert.Equal(t, "001", StringLeftPad("1", 3, "0"))
	assert.Equal(t, "****1", StringLeftPad("1", 5, "*"))
}

func TestStringDynamicReplacer(t *testing.T) {
	original := "Hello, {name}! Welcome to {place}."
	replaced := StringDynamicReplacer(original, "{name}", "Alice", "{place}", "Wonderland")
	expected := "Hello, Alice! Welcome to Wonderland."
	assert.Equal(t, expected, replaced)
}

func TestConvertAnyToJSONString(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    any
		Expected string
	}{
		{
			Name: "Struct input",
			Input: struct {
				Name string `json:"name"`
				Age  int    `json:"age"`
			}{Name: "Alice", Age: 30},
			Expected: `{"name":"Alice","age":30}`,
		},
		{Name: "Slice input", Input: []int{1, 2, 3}, Expected: `[1,2,3]`},
		{Name: "Nil input", Input: nil, Expected: `null`},
		{Name: "Unsupported input", Input: func() {}, Expected: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, ConvertAnyToJSONString(tc.Input))
		})
	}
}

func TestToString(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    any
		Expected string
	}{
		{"String input", "hello", "hello"},
		{"Integer input", 42, "42"},
		{"Float input", 3.14, "3.14"},
		{"Boolean input", true, "true"},
		{"Nil input", nil, "<nil>"},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := ToString(tc.Input)
			assert.Equal(t, tc.Expected, result)
		})
	}
}

func TestStringJoinSlice(t *testing.T) {
	testCases := []struct {
		Name      string
		Input     []any
		Separator string
		Expected  string
	}{
		{"Integer inputs", []any{1, 2, 3}, ",", "1,2,3"},
		{"String inputs", []any{"a", "b", "c"}, "|", "a|b|c"},
		{"Mix inputs", []any{"John", 25, true, 3.14}, " - ", "John - 25 - true - 3.14"},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := StringJoinSlice(tc.Input, tc.Separator)
			assert.Equal(t, tc.Expected, result)
		})
	}
}
