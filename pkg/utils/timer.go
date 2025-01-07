package utils

import "time"

func GetUTCTime() time.Time {
	return time.Now().UTC()
}
