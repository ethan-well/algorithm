package sortgo

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int64{12,121,789,1,3,10,44, 99, 33}
	QuickSort(arr, 0, len(arr) - 1)

	t.Logf("arr: %v", arr)
}

func QuickSort(arr []int64, l, r int){
	if l >= r {
		return
	}

	key := arr[l]
	i, j := l, r
	for i < j {
		for i < j && arr[j] > key {
			j --
		}

		if i < j {
			arr[i] = arr[j]
			i++
		}

		for i < j && arr[i] < key {
			i ++
		}
		if i < j {
			arr[j] = arr[i]
			j --
		}
	}

	arr[i] = key

	QuickSort(arr, l, i - 1)
	QuickSort(arr, i + 1, r)
}