package time

import "time"

func IsTimeBetween(t, start, end time.Time) bool {
	return (t.After(start) || t.Equal(start)) && (t.Before(end) || t.Equal(end))
}

func IsDateBetween(date, start, end string) bool {
	return date >= start && date <= end
}

func IsValidDateOnly(s string) bool {
	_, err := time.Parse(time.DateOnly, s)
	return err == nil
}
