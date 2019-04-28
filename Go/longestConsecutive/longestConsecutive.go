package longestconsecutive

func longestConsecutive(nums []int) int {
	numMap := make(map[int]int)
	maxLen := 0
	for _, num := range nums {
		tmp := numMap[num]

		rightOpt := num + 1
		for ; numMap[rightOpt] > 0; rightOpt++ {
			numMap[rightOpt]++
		}

		if rightCount := numMap[rightOpt-1]; rightCount > 0 {
			numMap[num] = rightCount
		} else {
			numMap[num] = 1
		}

		step := 0
		if tmp == 0 {
			step = numMap[num]
		} else {
			step = 1
		}

		var maxLen int
		leftOpt := num - 1
		for ; numMap[leftOpt] > 0; leftOpt-- {
			numMap[leftOpt] += step
		}

		if numMap[num-1] > numMap[num] {
			numMap[num] = numMap[num-1]
		}

		if step == 0 {
			for rightOpt--; rightOpt >= num; rightOpt-- {
				numMap[rightOpt] = numMap[num]
			}
		}

		for _, val := range nums {
			if numMap[val] > maxLen {
				maxLen = numMap[val]
			}
		}
	}

	for _, value := range numMap {
		if value > maxLen {
			maxLen = value
		}
	}
	return maxLen
}
