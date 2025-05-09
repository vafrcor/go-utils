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
