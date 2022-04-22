package roman

import (
	"testing"
)

func TestCode(t *testing.T) {
	str := "MCMXCIV"
	result := romanToInt(str, t)
	t.Logf("result: %d", result)
}

func romanToInt(s string, t *testing.T) int {
	RomanNumberMap := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	var result int
	for i := 0; i < len(s); {

		var c1 = ""
		var c2 = ""

		if len(s)-i-2 >= 0 {
			c1 = s[len(s)-i-2 : len(s)-i]
		}

		c2 = s[len(s)-i-1 : len(s)-i]

		t.Logf("i: %d, c1: %s, c2: %s", i, c1, c2)

		if num1, ok := RomanNumberMap[c1]; ok{
			result += num1
			i += 2
			continue
		}

		if num2, ok := RomanNumberMap[c2]; ok {
			result += num2
		}

		i += 1
	}

	return result
}
