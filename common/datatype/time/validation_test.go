package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsDateBetween(t *testing.T) {
	testCases := []struct {
		Name     string
		Date     string
		Start    string
		End      string
		Expected bool
	}{
		{"InRange", "2021-04-20", "2021-04-19", "2021-04-21", true},
		{"OutOfRangeBefore", "2021-04-18", "2021-04-19", "2021-04-21", false},
		{"OutOfRangeAfter", "2021-04-22", "2021-04-19", "2021-04-21", false},
		{"ExactStart", "2021-04-19", "2021-04-19", "2021-04-21", true},
		{"ExactEnd", "2021-04-21", "2021-04-19", "2021-04-21", true},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := IsDateBetween(tc.Date, tc.Start, tc.End)
			assert.Equal(t, tc.Expected, result)
		})
	}
}

func TestIsTimeBetween(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    time.Time
		Start    time.Time
		End      time.Time
		Expected bool
	}{
		{"InRange", time.Date(2021, 04, 20, 10, 0, 0, 0, time.UTC), time.Date(2021, 04, 19, 9, 0, 0, 0, time.UTC), time.Date(2021, 04, 21, 11, 0, 0, 0, time.UTC), true},
		{"OutOfRangeBefore", time.Date(2021, 04, 18, 10, 0, 0, 0, time.UTC), time.Date(2021, 04, 19, 9, 0, 0, 0, time.UTC), time.Date(2021, 04, 21, 11, 0, 0, 0, time.UTC), false},
		{"OutOfRangeAfter", time.Date(2021, 04, 22, 10, 0, 0, 0, time.UTC), time.Date(2021, 04, 19, 9, 0, 0, 0, time.UTC), time.Date(2021, 04, 21, 11, 0, 0, 0, time.UTC), false},
		{"ExactStart", time.Date(2021, 04, 19, 9, 0, 0, 0, time.UTC), time.Date(2021, 04, 19, 9, 0, 0, 0, time.UTC), time.Date(2021, 04, 21, 11, 0, 0, 0, time.UTC), true},
		{"ExactEnd", time.Date(2021, 04, 21, 11, 0, 0, 0, time.UTC), time.Date(2021, 04, 19, 9, 0, 0, 0, time.UTC), time.Date(2021, 04, 21, 11, 0, 0, 0, time.UTC), true},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := IsTimeBetween(tc.Input, tc.Start, tc.End)
			assert.Equal(t, tc.Expected, result)
		})
	}
}

func TestIsValidDateOnly(t *testing.T) {
	testCases := []struct {
		Name     string
		Date     string
		Expected bool
	}{
		{"valid", "2021-04-20", true},
		{"invalid_day", "2021-04-32", false},
		{"invalid_format", "202-04-22", false},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := IsValidDateOnly(tc.Date)
			assert.Equal(t, tc.Expected, result)
		})
	}
}
