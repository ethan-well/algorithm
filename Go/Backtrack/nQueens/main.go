package main

import "fmt"

var count int

func absEqual(a, b int) bool {
	result := (a == b) || (a+b == 0)
	return result
}

func isNotConflict(a *[8]int, k int) bool {
	for i := 0; i < k; i++ {
		for j := i + 1; j <= k; j++ {
			if a[i] == a[j] || absEqual(a[i]-a[j], i-j) {
				return false
			}
		}
	}
	return true
}

// 第 k 个元素
func backTrack(k int, a *[8]int) {
	// 本轮结束
	if k >= 8 {
		for i := 0; i < 8; i++ {
			fmt.Println(*a)
		}
		count++
	} else {
		// 每个元素都有 0 - 7 八种情况
		for i := 0; i < 8; i++ {
			a[k] = i
			if isNotConflict(a, k) {
				backTrack(k+1, a)
			}
		}
	}
}

func main() {
	var a [8]int
	backTrack(0, &a)
	fmt.Println(count)
}
