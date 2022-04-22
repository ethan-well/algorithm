package searchSameNumber

import (
	"fmt"
	"testing"
)

func SearchSameNumber(arr1, arr2 []int) []int {

	arr1NumberCount := make(map[int]int)

	numberCount := make(map[int]int)

	sameNumberMap := make(map[int]int)
	for i := 0; i < len(arr1); i ++ {

		arr1NumberCount[arr1[i]] ++

		arr2NumberCount := make(map[int]int)
		for j := 0; j < len(arr2); j++ {

			arr2NumberCount[arr2[j]] ++

			if arr1[i] == arr2[j] {

				numberCount[arr1[i]] = arr1NumberCount[arr1[i]]
				if arr1NumberCount[arr1[i]] > arr2NumberCount[arr1[i]] {
					numberCount[arr1[i]] = arr2NumberCount[arr1[i]]
				}

				sameNumberMap[i] = arr1[i]
			}
		}
	}

	var result []int
	for _, v := range sameNumberMap {
		if numberCount[v] > 0 {
			result = append(result, v)
		}
		numberCount[v] --
	}

	return result
}

func TestSearchSameNumber(t *testing.T) {
	arr1 := []int{2,2,2,4,4,5}
	arr2 := []int{1,3,2,2,4,4,5}
	arr := SearchSameNumber(arr1, arr2)
	fmt.Println(arr)
}
