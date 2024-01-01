package moni

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	remainC := 0
	profit := make([]int, 0)
	maxStep := 0
	nowStep := 0
	getStepProfit := func() {
		stepCustomer := 4
		if remainC < 4 {
			stepCustomer = remainC
		}
		remainC -= stepCustomer
		stepProfit := stepCustomer*boardingCost - runningCost
		if nowStep > 0 {
			newOne := stepProfit + profit[nowStep-1]
			profit = append(profit, newOne)
			nowStep++
			if stepProfit > 0 {
				maxStep = nowStep
			}
		} else {
			profit = append(profit, stepProfit)
			nowStep++
			maxStep = 1
		}
	}
	for _, customer := range customers {
		remainC += customer
		getStepProfit()
	}
	for remainC > 0 {
		getStepProfit()
	}
	if profit[maxStep-1] <= 0 {
		return -1
	}
	return maxStep
}
