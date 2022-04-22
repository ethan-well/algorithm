package sortgo

import "testing"

func TestSelectionSort(t *testing.T) {
	arr := SelectionSort([]int64{12,121,789,1,3,10,44, 99, 33})
	t.Logf("arr: %v", arr)
}

func SelectionSort(arr []int64) []int64 {
	var min int64
	var minIndex int
	for i := 0; i < len(arr); i ++ {
		min = arr[i]
		minIndex = i
		for j := i+1; j < len(arr); j ++ {
			if arr[j] < min {
				min = arr[j]
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}
