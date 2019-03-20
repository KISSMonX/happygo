package main

import (
	"fmt"
	"testing"
	"time"
)

// WeekStart 指定年份和周得出周开始日期时间
func WeekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}

// WeekRange 指定年份和周数得出起始时间
func WeekRange(year, week int) (start, end time.Time) {
	start = WeekStart(year, week)
	end = start.AddDate(0, 0, 6)
	return
}

func TestWeekNumber(t *testing.T) {
	y, w := time.Now().ISOWeek()
	fmt.Println(y, w)

	start, end := WeekRange(2019, 11)
	fmt.Println(start, end)
}
