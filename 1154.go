package main

import (
	"time"
)

func dayOfYear(date string) int {
	dateT, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}
	countDays := func(month int, run bool) int {
		switch month {
		case 1, 3, 5, 7, 8, 10, 12:
			return 31
		case 4, 6, 9, 11:
			return 30
		case 2:
			if run {
				return 29
			} else {
				return 28
			}
		default:
			return 0
		}
	}
	y, m, d := dateT.Year(), dateT.Month(), dateT.Day()
	// 400年一润，4年一润，百年不润
	run := y%400 == 0 || (y%4 == 0 && y%100 != 0)
	days := d
	for i := time.January; i < m; i++ {
		days += countDays(int(i), run)
	}
	return days
}
