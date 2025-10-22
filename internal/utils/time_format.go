package utils

import "time"

func FromStringToTime(input string) (time.Time, error) {
	layout := "2006-01-02"

	return time.Parse(layout, input)
}
