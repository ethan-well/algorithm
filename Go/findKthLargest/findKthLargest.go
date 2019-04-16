package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	quicksort(left, right, nums[:])

	return nums[k-1]
}

func quicksort(left, right int, arr []int) {
	if left > right {
		return
	}
	i, j := left, right
	for i < j {
		for i < j && arr[i] > arr[left] {
			i++
		}

		for i < j && arr[left] >= arr[j] {
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[left], arr[i] = arr[i], arr[left]

	quicksort(left, i-1, arr)
	quicksort(i+1, right, arr)
}

func main() {
	arr := [6]int{3, 2, 1, 5, 6, 4}

	result := findKthLargest(arr[:], 2)

	fmt.Println(result)
}
