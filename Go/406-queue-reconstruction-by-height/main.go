package main

func main() {
	people := [][]int{ {7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}

	reconstructQueue(people)
}

func reconstructQueue(people [][]int) [][]int {

	var heightQueue [][]int

	for _, p := range people {
		insertToQueue(&heightQueue, 0, len(heightQueue)-1, p)
	}

	var sortedQueue = make([][]int, len(people))
	for _, p := range heightQueue {

			var j = -1
			var i = 0
			for ; j < p[1]; i ++ {
				if sortedQueue[i] == nil ||
					sortedQueue[i][0] == p[0] {

					j ++
				}
			}

			if j == p[1] {
				sortedQueue[i-1] = []int{p[0], p[1]}
			}

	}

	return sortedQueue
}

func insertToQueue(heightQueue *[][]int, i, j int, p []int) {

	if len(*heightQueue) == 0 {
		*heightQueue = [][]int{p}
	}

	n := (i + j )/2

	if i >= j {
		if (*heightQueue)[n][0] < p[0] {
			arr := append([][]int{p},  (*heightQueue)[n+1:]...)
			*heightQueue = append((*heightQueue)[:n+1], arr...)
			return
		} else {
			arr := append([][]int{p}, (*heightQueue)[n:]...)
			*heightQueue = append((*heightQueue)[0:n], arr...)
			return
		}
	}

	if p[0] >= (*heightQueue)[n][0] {
		insertToQueue(heightQueue, n+1, j, p)
	}

	if p[0] < (*heightQueue)[n][0] {
		insertToQueue(heightQueue, i, n-1, p)
	}

	return
}