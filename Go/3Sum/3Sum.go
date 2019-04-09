package main

import (
	"fmt"
	"sort"
)

// func intSlice(arr []int) {
// 	sort.Ints(arr)
// }

func threeSum(arr []int) [][]int {
	sort.Ints(arr)
	var left, right int
	var result [][]int
	for i := 0; i < len(arr)-2 && arr[i] <= 0; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		left = i + 1
		right = len(arr) - 1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]

			switch {
			case sum == 0:
				tem := []int{arr[i], arr[left], arr[right]}
				result = append(result, tem)

				left++
				right--
				for left < right && arr[left] == arr[left-1] {
					left++
				}
				for left < right && arr[right] == arr[right+1] {
					right--
				}
			case sum < 0:
				left = left + 1
			default:
				right = right - 1
			}
		}
	}

	return result
}

func main() {
	var arr = [...]int{-1, 0, 1, 2, -1, -4}
	// intSlice(arr[:])
	result := threeSum(arr[:])
	fmt.Printf("result: %v\n", result)
}
