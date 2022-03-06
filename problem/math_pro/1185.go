package math_pro

/*
给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。

输入为三个整数：day、month 和year，分别表示日、月、年。

您返回的结果必须是这几个值中的一个{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}。

给出的日期一定是在 1971 到 2100 年之间的有效日期。
*/

/*
解答

题目保证日期是在 1971 到 2100 之间，我们可以计算给定日期距离 1970 的最后一天（星期四）间隔了多少天，从而得知给定日期是周几。

具体的，可以先通过循环处理计算年份在 [1971, year - 1][1971,year−1] 时间段，经过了多少天（注意平年为 365365，闰年为 366366）；然后再处理当前年 yearyear 的月份在 [1, month - 1][1,month−1] 时间段 ，经过了多少天（注意当天年是否为闰年，特殊处理 22 月份），最后计算当前月 monthmonth 经过了多少天，即再增加 dayday 天。

得到距离 1970 的最后一天（星期四）的天数后进行取模，即可映射回答案。

*/

var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

func dayOfTheWeek(day, month, year int) string {
	days := 0
	// 输入年份之前的年份的天数贡献
	days += 365*(year-1971) + (year-1969)/4
	// 输入年份中，输入月份之前的月份的天数贡献
	for _, d := range monthDays[:month-1] {
		days += d
	}
	if month >= 3 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		days++
	}
	// 输入月份中的天数贡献
	days += day
	return week[(days+3)%7]
}
