package main

func firstDayBeenInAllRooms(nextVisit []int) int {
	n := len(nextVisit)
	visitCount := make([]int, n)
	allVisit := func() bool {
		for _, v := range visitCount {
			if v == 0 {
				return false
			}
		}
		return true
	}
	visitCount[0] = 1
	lastVisit := 0
	allV := false
	nextV := 0
	days := 0
	mod := 1000000000 + 7
	for !allV {
		if count := visitCount[lastVisit]; count%2 != 0 {
			nextV = nextVisit[lastVisit]
		} else {
			nextV = (lastVisit + 1) % n
		}
		visitCount[nextV] = visitCount[nextV] + 1
		lastVisit = nextV
		allV = allVisit()
		days++
		if days >= mod {
			days = days % mod
		}
	}
	return days
}

func firstDayBeenInAllRooms(nextVisit []int) int {
	mod := 1000000007
	dp := make([]int, len(nextVisit))

	dp[0] = 2 //初始化原地待一天 + 访问下一个房间一天
	for i := 1; i < len(nextVisit); i++ {
		to := nextVisit[i]
		dp[i] = dp[i-1] + 2
		if to != 0 {
			dp[i] = (dp[i] - dp[to-1] + mod) % mod //避免负数
		}
		dp[i] = (dp[i] + dp[i-1]) % mod //题目保证n >= 2
	}

	return dp[len(nextVisit)-2]
}
