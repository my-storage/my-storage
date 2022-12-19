package time

import "time"

func GetCurrentTime() int {
	result := time.Now().UnixNano() / int64(time.Millisecond)
	return int(result)
}

func DaysToMilliseconds(days int) int {
	daysInMilliseconds := days * 24 * 60 * 60 * 1000

	return daysInMilliseconds
}

func HoursToMilliseconds(hours int) int {
	hoursInMilliseconds := hours * 60 * 60 * 1000

	return hoursInMilliseconds
}

func MinutesToMilliseconds(minutes int) int {
	minutesInMilliseconds := minutes * 60 * 1000

	return minutesInMilliseconds
}

func SecondsToMilliseconds(seconds int) int {
	secondsInMilliseconds := seconds * 1000

	return secondsInMilliseconds
}

func MillisecondsToTime(milliseconds int) time.Time {
	result := time.Unix(0, int64(milliseconds)*int64(time.Millisecond))
	return result
}

func TimeToMilliseconds(date time.Time) int {
	result := date.UnixNano() / int64(time.Millisecond)
	return int(result)
}
