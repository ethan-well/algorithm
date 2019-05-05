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

		rightCount := 0
		for rightOpt := num + 1; numMap[rightOpt] > 0; rightOpt++ {
			rightCount++
		}

		leftCount := 0
		for leftOpt := num - 1; numMap[leftOpt] > 0; leftOpt-- {
			leftCount++
		}

		for rightOpt := num + rightCount; numMap[rightOpt] > 0; rightOpt-- {
			numMap[rightOpt] = rightCount + leftCount + 1
		}
	}

	for _, value := range numMap {
		if value > maxLen {
			maxLen = value
		}
	}
	return maxLen
}
