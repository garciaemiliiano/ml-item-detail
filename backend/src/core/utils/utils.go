package utils

import "time"

func ParseDateStr(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		return time.Time{}
	}
	return t
}
