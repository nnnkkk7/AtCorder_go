package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s []int = make([]int, n)
	for i := 0; i < n; i++ {
		var t int
		fmt.Scan(&t)
		s[i] = t
	}
	var dp []int = make([]int, n)
	for i := 1; i < n; i++ {
		if i == 1 {
			dp[i] = abs(s[i] - s[i-1])
		} else {
			s1 := dp[i-1] + abs(s[i]-s[i-1])
			s2 := dp[i-2] + abs(s[i]-s[i-2])
			if s1 >= s2 {
				dp[i] = s2
			} else {
				dp[i] = s1
			}
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
