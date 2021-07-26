package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	var s []int = make([]int, n)
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		s[i] = t
	}
	var dp []int = make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		if i == 1 {
			dp[i] = abs(s[i] - s[i-1])
		} else {
			var ss []int
			if i <= k {
				for j := 1; j <= i; j++ {
					ss = append(ss, dp[i-j]+abs(s[i]-s[i-j]))
				}
			} else {
				for j := 1; j <= k; j++ {
					ss = append(ss, dp[i-j]+abs(s[i]-s[i-j]))
				}
			}
			dp[i] = min(ss)
		}
	}
	fmt.Println(dp[n-1])
}

// 絶対値を返す
func abs(num int) int {
	if num < 0 {
		num *= -1
	}
	return num
}

//スライスの最小値を返す
func min(num []int) int {
	var min int = 1000000000
	for i := 0; i < len(num); i++ {
		if num[i] < min {
			min = num[i]
		}
	}
	return min
}
