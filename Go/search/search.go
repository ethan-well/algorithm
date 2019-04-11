package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return searchTarget(nums, 0, len(nums)-1, target)
}

func searchTarget(nums []int, left, right, target int) int {
	targetIndex := -1

	if left > right || left < 0 || right < 0 || right >= len(nums) || left >= len(nums) {
		return targetIndex
	}

	m := (left + right) / 2

	if nums[m] == target {
		targetIndex = m
		return targetIndex
	}

	if nums[left] < nums[m] && nums[left] <= target && target < nums[m] {
		targetIndex = searchTarget(nums, left, m-1, target)
		return targetIndex
	}

	if nums[m] < nums[right] && nums[m] < target && target <= nums[right] {
		targetIndex = searchTarget(nums, m+1, right, target)
		return targetIndex
	}

	if nums[left] > nums[m] || nums[m] > nums[right] {
		// search first halft
		targetIndex = searchTarget(nums, left, m-1, target)
		// search second halft
		if targetIndex == -1 {
			targetIndex = searchTarget(nums, m+1, right, target)
		}
	}

	return targetIndex
}

func main() {
	arr := []int{2, 3, 4, 5, 1}
	target := 1
	fmt.Println(search(arr, target))
}
