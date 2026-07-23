package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetDiffDayBetweenTime001(t *testing.T) {
	start, _ := time.Parse("2006-01-02", "2021-01-01")
	end := start.AddDate(0, 0, 30)
	assert.Equal(t, 30, GetDiffDaysBetweenTime(start, end))
}

func TestGetDiffDayBetweenTime002(t *testing.T) {
	start, _ := time.Parse("2006-01-02", "2021-01-01")
	end := start.AddDate(0, 0, -30)
	assert.Equal(t, 30, GetDiffDaysBetweenTime(start, end))
}

func TestConvertCustomTzDateToUtc001(t *testing.T) {
	convert, err := ConvertDateWithCustomTzIntoUtcTime("2021-04-21", "Asia/Jakarta", "2006-01-02")
	assert.Equal(t, "2021-04-20 17:00:00 UTC", convert.Format("2006-01-02 15:04:05 MST"))
	assert.NoError(t, err)
}

func TestConvertCustomTzDateToUtc002(t *testing.T) {
	convert, err := ConvertDateWithCustomTzIntoUtcTime("2021-04-21", "America/New_York", "2006-01-02")
	assert.Equal(t, "2021-04-21 04:00:00 UTC", convert.Format("2006-01-02 15:04:05 MST"))
	assert.NoError(t, err)
}

func TestConvertCustomTzDateToUtc003(t *testing.T) {
	sourceDate := time.Date(2021, 04, 22, 0, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60))
	convert, err := ConvertDateWithCustomTzIntoUtcTime(sourceDate.Format(time.RFC3339), "Asia/Jakarta", time.RFC3339)
	assert.Equal(t, "2021-04-21 17:00:00 UTC", convert.Format("2006-01-02 15:04:05 MST"))
	assert.NoError(t, err)
}

func TestConvertUtcRFC3339DatetimeIntoCustomDateTz001(t *testing.T) {
	convert, err := ConvertUtcRFC3339DatetimeIntoCustomDateTz("2021-04-20T17:00:00Z", "Asia/Jakarta")
	assert.Equal(t, "2021-04-21 00:00:00 WIB", convert.Format("2006-01-02 15:04:05 MST"))
	assert.NoError(t, err)
}

func TestConvertUtcRFC3339DatetimeIntoCustomDateTz002(t *testing.T) {
	convert, err := ConvertUtcRFC3339DatetimeIntoCustomDateTz("2021-04-20T02:00:00Z", "Asia/Jakarta")
	assert.Equal(t, "2021-04-20 09:00:00 WIB", convert.Format("2006-01-02 15:04:05 MST"))
	assert.NoError(t, err)
}

func TestConvertUtcRFC3339DatetimeIntoCustomDateTz003(t *testing.T) {
	convert, err := ConvertUtcRFC3339DatetimeIntoCustomDateTz("2021-04-20T02:00:00Z", "America/New_York")
	assert.Equal(t, "2021-04-19 22:00:00 EDT", convert.Format("2006-01-02 15:04:05 MST"))
	assert.NoError(t, err)
}

func TestGetTimeFromStringWithFormat(t *testing.T) {
	x, err := GetTimeFromStringWithFormat("2021-04-20", "2006-01-02")
	assert.NoError(t, err)
	assert.Equal(t, time.Date(2021, 04, 20, 0, 0, 0, 0, time.UTC), *x)
}

func TestGetDurationStringFromTotalSeconds(t *testing.T) {
	assert.Equal(t, "00:00:50", GetDurationStringFromTotalSeconds(50))
	assert.Equal(t, "00:02:10", GetDurationStringFromTotalSeconds(130))
	assert.Equal(t, "02:50:40", GetDurationStringFromTotalSeconds(10240))
}

func TestGetDateStringFromCustomFormat(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    string
		Format   string
		Expected string
		HasError bool
	}{
		{"ValidDate", "2026-03-09", time.DateOnly, "2026-03-09", false},
		{"ValidDateTime", "2026-03-10T10:30:59Z", time.RFC3339, "2026-03-10", false},
		{"ValidDateTime", "2026-02-30T10:30:59Z", time.RFC3339, "2026-03-10", true},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := GetDateOnlyStringFromCustomFormat(tc.Input, tc.Format)
			if tc.HasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, result)
			}
		})
	}
}

func TestConvertHHMMWithCustomTzIntoUtcTime(t *testing.T) {
	testCases := []struct {
		Name     string
		HHMM     string
		Timezone string
		Expected string
		HasError bool
	}{
		{"Jakarta UTC+7 afternoon", "13:00", "Asia/Jakarta", "06:00", false},
		{"Jakarta UTC+7 midnight wrap", "01:00", "Asia/Jakarta", "18:00", false},
		{"Kolkata UTC+5:30", "14:00", "Asia/Kolkata", "08:30", false},
		{"UTC no change", "09:30", "UTC", "09:30", false},
		{"invalid timezone", "13:00", "Invalid/Zone", "", true},
		{"invalid HHMM", "99:99", "Asia/Jakarta", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := ConvertHHMMWithCustomTzIntoUtcTime(tc.HHMM, tc.Timezone)
			if tc.HasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, result)
			}
		})
	}
}
