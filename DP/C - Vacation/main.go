//FIXME
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		var t []int = make([]int, 3)
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		t[0] = a
		t[1] = b
		t[2] = c
		s[i] = t
	}
	var dp []int = make([]int, n)
	var inlist []int = make([]int, n)
	var index int
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i], index = max(s[i])
			inlist[i] = index
		} else {
			dp[i], index = max(s[i])
			if inlist[i-1] == index {
				s[i][index] = -1
				dp[i], index = max(s[i])
			}
			inlist[i] = index
		}
	}
	fmt.Println(sum(dp))
}

//スライスの最大値と何番目かを返す
func max(num []int) (int, int) {
	var max int = -10
	var index int
	for i := 0; i < len(num); i++ {
		if num[i] > max {
			max = num[i]
			index = i
		}
	}
	return max, index
}

// スライスの要素の合計
func sum(num []int) int {
	var ans int
	for i := 0; i < len(num); i++ {
		ans += num[i]
	}
	return ans
}
