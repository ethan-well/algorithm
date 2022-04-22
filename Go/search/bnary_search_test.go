package search

import (
	"testing"
)

func TestSomeCode(t *testing.T) {
   arr := []int64{0,1,1, 10, 88, 90, 100}
   i := binarySearch(arr, 90)

   t.Logf("i: %d", i)
}

func binarySearch(arr []int64, target int64) int {
	l := 0
	r := len(arr) - 1

	for l < r {
		mid := l + (r - 1) / 2

		if target == arr[mid] {
			return mid
		}

		if target < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
