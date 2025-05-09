package time

import (
	"math"
	"time"

	"github.com/vafrcor/go-utils/env"
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
	durX := env.GetEnvValue(envDurStr, "string", defaultDurStr).(string)
	return ConvertDurationStringToTimeDuration(durX, defaultDurStr)
}
