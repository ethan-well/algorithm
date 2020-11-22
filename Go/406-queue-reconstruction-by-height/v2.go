package main

import "log"

func reconstructQueueV2(people [][]int) [][]int {

	quickSortV2(&people, 0, len(people)-1)

	log.Printf("people: %v", people)

	var sortedQueue = make([][]int, len(people))
	//
	//for _, p := range people {
	//
	//	var j = -1
	//	var i = 0
	//	for ; j < p[1]; i ++ {
	//		if sortedQueue[i] == nil ||
	//			sortedQueue[i][0] == p[0] {
	//
	//			j ++
	//		}
	//	}
	//
	//	if j == p[1] {
	//		sortedQueue[i-1] = p
	//	}
	//}

	return sortedQueue
}

func quickSort(arr *[]int, begin, end int) {

	if begin < end {
		tmp := (*arr)[begin]
		i := begin
		j := end

		for  i < j {
			for i < j && (*arr)[j] > tmp {
				j --
			}

			if i < j {
				(*arr)[i] = (*arr)[j]
				i ++
			}

			for i < j && (*arr)[i] < tmp {
				i ++
			}

			if i < j {
				(*arr)[j] =  (*arr)[i]
				j--
			}
		}

		(*arr)[i] = tmp
		quickSort(arr, 0, i-1)
		quickSort(arr, i+1, end)
	}
}

func quickSortV2(arr *[][]int, begin, end int) {

	if begin < end {
		tmp := (*arr)[begin]
		i := begin
		j := end

		for  i < j {
			for i < j && (*arr)[j][0] > tmp[0] {
				j --
			}

			if i < j {
				(*arr)[i] = (*arr)[j]
				i ++
			}

			for i < j && (*arr)[i][0] < tmp[0] {
				i ++
			}

			if i < j {
				(*arr)[j] =  (*arr)[i]
				j--
			}
		}

		(*arr)[i] = tmp
		quickSortV2(arr, 0, i-1)
		quickSortV2(arr, i+1, end)
	}
}