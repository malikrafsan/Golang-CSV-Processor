package utils

import "time"

func IdFromTime() (string, error) {
	return time.Now().Format("20060102150405"), nil
}
