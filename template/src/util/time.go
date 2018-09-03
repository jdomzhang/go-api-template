package util

import "time"

// GetDate will return the date without time
func GetDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetWeekBeginDate will return monday of the given time
func GetWeekBeginDate(t time.Time) time.Time {
	t = GetDate(t)
	weekday := t.Weekday()
	days := 0
	switch weekday {
	case time.Sunday:
		days = 7 + int(weekday) - int(time.Monday)
	default:
		days = int(weekday) - int(time.Monday)
	}

	return t.AddDate(0, 0, -1*days)
}

// GetMonthBeginDate will return monday of the given time
func GetMonthBeginDate(t time.Time) time.Time {
	t = GetDate(t)
	dayofmonth := t.Day()

	return t.AddDate(0, 0, -1*dayofmonth+1)
}
