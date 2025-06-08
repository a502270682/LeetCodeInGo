package dp

func LongestPalindrome(s string) string {
	length := len(s)
	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, length)
	}
	// init
	for i := 0; i < length; i++ {
		dp[i][i] = true
	}
	start := 0
	maxLength := 1
	for j := 1; j < length; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				continue
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > maxLength {
				maxLength = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+maxLength]
}
