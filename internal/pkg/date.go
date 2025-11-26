package pkg

import "time"

func GetCurrentDateTime() time.Time {
	return time.Now()
}

func HourToNumber(hours int) int64 {
	return int64(hours) * int64(time.Hour)
}
