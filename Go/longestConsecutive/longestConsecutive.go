package longestconsecutive

func longestConsecutive(nums []int) int {
	numMap := make(map[int]int)
	maxLen := 0
	for _, num := range nums {
		if numMap[num] != 0 {
			continue
		} else {
			numMap[num] = 1
		}

		for rightOpt := num + 1; numMap[rightOpt] > 0; rightOpt++ {
			numMap[rightOpt]++
		}

		if rightCount := numMap[num+1]; rightCount > numMap[num] {
			numMap[num] = rightCount
		}

		step := numMap[num]

		for leftOpt := num - 1; numMap[leftOpt] > 0; leftOpt-- {
			numMap[leftOpt] += step
		}

		if numMap[num-1] > numMap[num] {
			numMap[num] = numMap[num-1]
		}

		for rightOpt := num + 1; numMap[rightOpt] > 0; rightOpt++ {
			numMap[rightOpt] = numMap[num]
		}
	}

	for _, value := range numMap {
		if value > maxLen {
			maxLen = value
		}
	}
	return maxLen
}
