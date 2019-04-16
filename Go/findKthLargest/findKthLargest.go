package main

import "fmt"

func findKthLargest(arr []int, k int) int {
	arrUniqed := arrayUniqByMap(arr[:])

	left, right := 0, len(arrUniqed)-1
	quicksort(left, right, arrUniqed[:])
	return arrUniqed[k-1]
}

func quicksort(left, right int, arr []int) {
	if left > right {
		return
	}
	i, j := left, right
	for i < j {
		for i < j && arr[left] <= arr[j] {
			j--
		}

		for i < j && arr[i] <= arr[left] {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[left], arr[i] = arr[i], arr[left]

	quicksort(left, i-1, arr)
	quicksort(i+1, right, arr)
}

func arrayUniqByMap(arr []int) []int {
	arrToMap := make(map[int]int)
	for _, value := range arr {
		arrToMap[value] = value
	}

	var arrUniqed []int
	for _, value := range arrToMap {
		arrUniqed = append(arrUniqed, value)
	}

	return arrUniqed
}

func main() {
	arr := [9]int{3, 2, 3, 1, 2, 4, 5, 5, 6}

	result := findKthLargest(arr[:], 4)

	fmt.Println(result)
}
