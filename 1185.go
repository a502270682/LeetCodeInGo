package main

var week = []string{"Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday"}
var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

func dayOfTheWeek(day, month, year int) string {

	days := 0
	// 输入年份之前的年份的天数贡献
	// 其中1968是润年，所以以其为基准，判断是否有润年需要+1天
	days += 365*(year-1971) + (year-1-1968)/4
	// 输入年份中，输入月份之前的月份的天数贡献
	for _, d := range monthDays[:month-1] {
		days += d
	}
	if month >= 3 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		//如果超过2月份，考虑是否有润年
		days++
	}
	// 输入当月的天数贡献
	days += day
	// 1970 年 1 月 1 日是星期四
	return week[days%7]
}
