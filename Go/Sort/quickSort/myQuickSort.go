package main

// func myQuickSort(left, right int, arr *[8000000]int) {
// 	l, r := left, right
// 	postion := arr[(left+right)/2]
// 	tmp := 0
// 	for l < r {
// 		for arr[l] < arr[postion] {
// 			l++
// 		}

// 		for arr[r] > arr[postion] {
// 			r--
// 		}

// 		if l > r {
// 			break
// 		}
// 	}
// 	tmp = arr[l]
// 	arr[l] = arr[r]
// 	arr[r] = tmp

// 	if l == r {
// 		l--
// 		r++
// 	}

// 	if left < l {
// 		myQuickSort(left, l, arr)
// 	}

// 	if right > r {
// 		myQuickSort(r, right, arr)
// 	}
// }
