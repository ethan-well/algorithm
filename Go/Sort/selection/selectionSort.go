package main

import "fmt"

func selectionSort(arr *[10]int) {
	var tmp int
	for i := 0; i < len(arr); i++ {
		tmp = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				tmp = j
			}
		}
		arr[i], arr[tmp] = arr[tmp], arr[i]
	}
}

func main() {
	arr := [10]int{4, 6, 1, 9, 7, 5, 20, 14, 13, 17}
	selectionSort(&arr)
	fmt.Println(arr)
}
