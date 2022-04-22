package sortgo

import "testing"

func TestInsertSort(t *testing.T) {
	arr := InsertSort([]int64{12,121,789,1,3,10,44, 99, 33})
	t.Logf("arr: %v", arr)
}


func InsertSort(arr []int64) []int64 {
	for i := 1; i < len(arr); i ++ {
		for j := i; j > 0 ; j -- {
			if arr[j] < arr[j - 1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}

	return arr
}