package utils

import (
	"time"
)

const (
	dateLayout = "1990-01-01T13:00:00:00Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowFormatted() string {
	return GetNow().Format(dateLayout)
}
