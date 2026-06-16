package main

import "fmt"

func waysToClimb(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1
	dp[3] = 2

	for i := 4; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-3]
	}
	return dp[n]
}

func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	fmt.Println(waysToClimb(n))
}
