package main

import (
	"testing"
)

func TestInsertToQueue(t *testing.T) {
	// var arr []int

	arr := []int{4, 5, 7, 7}

	t.Logf("arr length: %d", len(arr))

	insertToQueue(&arr, 0, len(arr) - 1, 6)

	t.Logf("arr: %v", arr)
}

func TestReconstructQueue(t *testing.T) {
	people := [][]int{ {7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}

	queue := reconstructQueue(people)
	t.Logf("queue: %v", queue)
}