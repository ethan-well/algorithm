package main

import "fmt"

func insertionSort(a *[10]int) {
	for i := 0; i < len(a); i++ {
		for j := i - 1; j >= 0; j-- {
			if a[j+1] < a[j] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
	}
}

func main() {
	arr := [10]int{4, 6, 1, 9, 7, 5, 20, 14, 13, 17}
	insertionSort(&arr)
	fmt.Println(arr)
}
