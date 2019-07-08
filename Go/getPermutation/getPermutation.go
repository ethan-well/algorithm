package main

func main() {
	getPermutation(3, 2)
	// seriesCount(3)
}

func getPermutation(n int, k int) []int {
	result_arr := make([]int, n, n)
	nums := make([]int, n, n)

	for index, _ := range nums {
		nums[index] = index + 1
	}

	seriesCount(3)
	// var seriesCount int
	// for i := n; i > 0; i-- {
	// 	seriesCount = seriesCount(i - 1)
	// 	fmt.Println("seriesCount", seriesCount)
	// 	startNumber = k / seriesCount
	// 	series = k % seriesCount
	// 	if series == 0 {
	// 		k = seriesCount
	// 	} else {
	// 		startNumber++
	// 	}

	// 	selectNumber(startNumber, result_arr, nums, n, i)
	// }

	return result_arr
}

// func selectNumber(startNumber, result_arr, nums) {
// 	start_index = startNumber - 1
// 	result_arr[n-i] = nums[start_index]
// 	nums = append(nums[:start_index], nums[startNumber:])
// }

func seriesCount(n int) int {
	if n <= 1 {
		return 1
	}

	return (n) * seriesCount(n-1)
}
