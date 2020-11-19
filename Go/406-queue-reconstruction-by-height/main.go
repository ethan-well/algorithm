package main

import "log"

func main() {
	people := [][]int{ {7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}

	reconstructQueue(people)
}

func reconstructQueue(people [][]int) [][]int {

	heightIMap := make(map[int][]int)

	var heightQueue []int

	for _, p := range people {
		heightIMap[p[0]] = append(heightIMap[p[0]], p[1])
		insertToQueue(&heightQueue, 0, len(heightQueue)-1, p[0])
	}

	for _, h := range heightQueue {
		nums := heightIMap[h]
		var numsSorted []int
		for _, n := range nums {
			insertToQueue(&numsSorted, 0, len(numsSorted) -1, n)
		}

		heightIMap[h] = numsSorted

		log.Printf("h: %d, numsSorted: %v", h, numsSorted)
	}

	var sortedQueue = make([][]int, len(people))
	for _, h := range heightQueue {
		for _, n := range heightIMap[h] {

			var j = -1
			var i = 0
			for ; j < n; i ++ {
				if sortedQueue[i] == nil ||
					sortedQueue[i][0] == h {

					j ++
				}
			}

			if j == n {
				sortedQueue[i-1] = []int{h, n}
			}
		}
	}

	return sortedQueue
}

func insertToQueue(heightQueue *[]int, i, j, h int) {

	if len(*heightQueue) == 0 {
		*heightQueue = []int{h}
		return
	}

	n := (i + j )/2

	if h == (*heightQueue)[n] {
		return
	}

	if i >= j {
		if (*heightQueue)[n] < h {
			arr := append([]int{h},  (*heightQueue)[n+1:]...)
			*heightQueue = append((*heightQueue)[:n+1], arr...)
			return
		} else {
			arr := append([]int{h}, (*heightQueue)[n:]...)
			*heightQueue = append((*heightQueue)[0:n], arr...)
			return
		}
	}

	if h > (*heightQueue)[n] {
		insertToQueue(heightQueue, n+1, j, h)
	}

	if h < (*heightQueue)[n]  {
		insertToQueue(heightQueue, i, n-1, h)
	}

	return
}