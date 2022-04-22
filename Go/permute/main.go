package main

import "fmt"

func main() {
	nums := []int{1,2,3}
	arrs := permute(nums)
	for _, arr := range arrs {
		fmt.Println(arr)
	}
}

func permute(nums []int) [][]int {

	var results [][]int

	for i := 0; i < len(nums); i ++ {

		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		if i + 1 == len(nums) {
			nums2 = append(nums2[0:i])
		} else {
			nums2 = append(nums2[0:i], nums2[i+1:]...)
		}

		arrs := permute(nums2)
		if len(arrs) == 0 {
			return append(results, []int{nums[i]})
		}

		for _, arr := range arrs {
			results = append(results, append([]int{nums[i]}, arr...))
		}

	}

	return results
}
