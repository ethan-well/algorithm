package longestconsecutive

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	nums := [6]int{100, 4, 200, 1, 3, 2}
	if result := longestConsecutive(nums[:]); result != 4 {
		t.Errorf("longestConsecutive() get %d, expected 4", result)
	}
}
