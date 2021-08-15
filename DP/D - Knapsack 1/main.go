package main

import (
	"fmt"
)

func main() {
	var n, w int
	fmt.Scan(&n, &w)
	wlist := make([]int, n)
	vlist := make([]int, n)
	for i := 0; i < n; i++ {
		var ws, v int
		fmt.Scan(&ws, &v)
		wlist[i] = ws
		vlist[i] = v
	}
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= w; j++ {
			if j-wlist[i] >= 0 {
				if dp[i][j-wlist[i]]+vlist[i] >= dp[i+1][j] {
					dp[i+1][j] = dp[i][j-wlist[i]] + vlist[i]
				}
			}
			if dp[i][j] >= dp[i+1][j] {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
	fmt.Println(dp[n][w])
}
