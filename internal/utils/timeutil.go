package utils

import (
	"math"
	"time"
)

func BetweenMinuit(aMilli int64, bMilli int64) int {
	aTime := time.UnixMilli(aMilli)
	bTime := time.UnixMilli(bMilli)

	duration := aTime.Sub(bTime)

	diffMinutes := int(math.Ceil(math.Abs(duration.Minutes())))
	return diffMinutes
}

func NowMilliSecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func Add(timestamp int64, duration time.Duration) int64 {
	t := time.UnixMilli(timestamp)
	newTime := t.Add(duration)
	return newTime.UnixMilli()
}

func Minus(timestamp int64, duration time.Duration) int64 {
	t := time.UnixMilli(timestamp)
	newTime := t.Add(-duration)
	return newTime.UnixMilli()
}
