package main

import "testing"

func TestInsertToQueue(t *testing.T) {
	arr := []int{1, 2, 4, 5, 8, 10}
	arr2 := insertToQueue(arr, 0)

	t.Logf("arr: %v", arr2)
}

func TestReconstructQueue(t *testing.T) {
	people := [][]int{ {7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}

	queue := reconstructQueue(people)
	t.Logf("queue: %v", queue)
}