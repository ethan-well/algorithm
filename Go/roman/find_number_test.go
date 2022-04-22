package roman

import (
	"sort"
	"testing"
)

func TestSomeCode(t *testing.T) {
	result := fourSum([]int{1,0,-1,0,-2,2}, 0)
	t.Logf("%v", result)
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)

	result := make([][]int, 0)
	for i := 0; i < len(nums) - 3;{

		for j := i + 1; j < len(nums) - 2; {

			for k, t := j + 1, len(nums) - 1; k < t; {
				sum := nums[i] + nums[j] + nums[k] + nums[t]

				if sum < target {
					k ++
					continue
				}


				if sum > target {
					t --
					continue
				}

				result = append(result, []int{nums[i], nums[j],nums[k],nums[t]})

				kTmp := k
				k ++
				for nums[k] == nums[kTmp] && k < t {
					k ++
				}

				tTmp := t
				t --
				for nums[t] == nums[tTmp] && t > j {
					t --
				}
			}

			jTmp := j
			j ++
			for nums[j] == nums[jTmp] && j < len(nums) - 2 {
				j ++
			}
		}

		iTmp := i
		i ++
		for nums[i] == nums[iTmp] && i < len(nums) - 3 {
			i++
		}
	}

	return result
}
