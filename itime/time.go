package itime

import "time"

func HourTimestamp(hour int) (timestamp int64) {
	now := time.Now()
	timestamp = time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, now.Location()).Unix()
	return
}
