package time

import (
	"errors"
	"fmt"
	"math"
	"time"

	system "github.com/vafrcor/go-utils/common/system"
)

func GetDiffDaysBetweenTime(start time.Time, end time.Time) int {
	duration := end.Sub(start)
	durationDays := int(duration.Hours() / 24)
	return int(math.Abs(float64(durationDays)))
}

func ConvertDurationStringToTimeDuration(durStr string, defaultDurStr string) time.Duration {
	duration, err := time.ParseDuration(durStr)
	if err != nil {
		duration, _ = time.ParseDuration(defaultDurStr)
	}
	return duration
}

func ConvertEnvDurationStringToTimeDuration(envDurStr string, defaultDurStr string) time.Duration {
	durX := system.GetEnvValue(envDurStr, "string", defaultDurStr).(string)
	return ConvertDurationStringToTimeDuration(durX, defaultDurStr)
}

func ConvertDateWithCustomTzIntoUtcTime(inputDateStr string, inputTz string, timeFormat string) (*time.Time, error) {
	loc, err := time.LoadLocation(inputTz)
	if err != nil {
		return nil, err
	}
	inputDate, err := time.Parse(timeFormat, inputDateStr)
	if err != nil {
		return nil, err
	}
	inputDateLocalTz := time.Date(inputDate.Year(), inputDate.Month(), inputDate.Day(), inputDate.Hour(), inputDate.Minute(), inputDate.Second(), 0, loc)
	dateInUtcTz := inputDateLocalTz.In(time.UTC)
	return &dateInUtcTz, nil
}

func ConvertUtcRFC3339DatetimeIntoCustomDateTz(utcDatetimeStr string, targetTz string) (*time.Time, error) {
	utcDate, err := time.Parse(time.RFC3339, utcDatetimeStr)
	if err != nil {
		return nil, err
	}
	loc, err := time.LoadLocation(targetTz)
	if err != nil {
		return nil, err
	}

	dateInCustomTz := utcDate.In(loc)
	return &dateInCustomTz, nil
}

func GetTimeFromStringWithFormat(dateStr string, timeFormat string) (*time.Time, error) {
	if dateStr == "" {
		return nil, errors.New("dateStr is empty")
	}
	// "2006-01-02"
	t, err := time.Parse(timeFormat, dateStr)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func GetDurationStringFromTotalSeconds(duration uint64) string {
	hours := duration / 3600
	minutes := (duration % 3600) / 60
	seconds := duration % 60
	return fmt.Sprintf(
		"%02d:%02d:%02d", hours, minutes, seconds,
	)
}

func GetUtcTime() time.Time {
	return time.Now().UTC()
}

func ConvertHHMMWithCustomTzIntoUtcTime(hhMM string, orgTimezone string) (string, error) {
	loc, err := time.LoadLocation(orgTimezone)
	if err != nil {
		return "", err
	}
	now := time.Now().In(loc)
	parsed, err := time.ParseInLocation("2006-01-02 15:04", now.Format("2006-01-02")+" "+hhMM, loc)
	if err != nil {
		return "", fmt.Errorf("invalid HH:MM format: %w", err)
	}
	return parsed.UTC().Format("15:04"), nil
}

func GetDateOnlyStringFromCustomFormat(timeStr string, format string) (string, error) {
	t, err := time.Parse(format, timeStr)
	if err != nil {
		return "", err
	}
	return t.Format(time.DateOnly), nil
}
