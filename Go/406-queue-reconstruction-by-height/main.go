package main

func main() {
	people := [][]int{ {7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}

	reconstructQueue(people)
}

func reconstructQueue(people [][]int) [][]int {

	heightIMap := make(map[int][]int)
	var heightQueue []int
	
	for _, p := range people {
		heightIMap[p[0]] = append(heightIMap[p[0]], p[1])
		heightQueue = insertToQueue(heightQueue, p[0])
	}

	// i: 身高从小到大的排序
	// h: 身高
	for _, h := range heightQueue {
		nums := heightIMap[h]
		var numsSorted []int
		for _, n := range nums {
			numsSorted = insertToQueue(numsSorted, n)
		}

		heightIMap[h] = numsSorted
	}

	var sortedQueue = make([][]int, len(people))
	for _, h := range heightQueue {
		for _, n := range heightIMap[h] {
			j := -1

			for i, p := range sortedQueue {

				if p == nil || p[0] == h {
					j ++
				}

				if j == n {
					sortedQueue[i] = []int{h, n}
					break
				}

			}
		}
	}

	return sortedQueue
}

func insertToQueue(heightQueue []int, h int) []int {
	for i, height := range heightQueue {
		if height == h {
			return heightQueue
		}
		if height > h {
			heightQueue2 := append([]int{h}, heightQueue[i:]...)
			heightQueue2 = append(heightQueue[0:i], heightQueue2...)

			return heightQueue2
		}

	}

	heightQueue = append(heightQueue, h)

	return heightQueue
}