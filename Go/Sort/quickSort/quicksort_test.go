package main

import "testing"

func TestQuickSort(t *testing.T) {

}

func TestQuickSort2(t *testing.T) {
	arr := []int{-9, 78, 0, 23, -567, 70, 123, 90, -23}

	QuickSort2(arr, 0, len(arr)-1)

}
