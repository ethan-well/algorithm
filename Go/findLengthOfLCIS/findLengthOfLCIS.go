package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	if arrLen := len(nums); arrLen <= 1 {
		return arrLen
	}
	maxLen := 1
	subLen := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			subLen = subLen + 1
		} else {
			subLen = 1
		}

		if subLen > maxLen {
			maxLen = subLen
		}

	}

	return maxLen
}

func main() {
	var arr = []int{1, 3, 5, 4, 7}

	len := findLengthOfLCIS(arr)
	fmt.Println(len)
}
