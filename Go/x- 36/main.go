package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum("1b", "2x"))
}

func sum(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	currentLen := len1
	if len1 < len2 {
		currentLen = len2
	}

	var result []uint8
	var n uint8
	for i := 0; i < currentLen; i ++ {
		var x, y uint8

		if i + 1 > len(str1) {
			x = '0'
		} else {
			x = str1[len(str1) - i - 1]
		}

		if i + 1 > len(str2) {
			y = '0'
		} else {
			y = str2[len(str2) - i - 1]
		}

		if x > '9' {
			x = x - 'a' + 10
		} else {
			x = x - '0'
		}

		if y > '9' {
			y = y - 'a' + 10
		} else {
			y = y - '0'
		}

		if x + y + n > 36 {
			result = append(result, x + y + n - 36)
			n = 1
		} else {
			result = append(result, x + y + n)
			n = 0
		}
	}

	if n != 0 {
		result = append(result, n)
	}

	var resultStr string
	for i := len(result) - 1; i >= 0; i -- {
		if result[i] > 9 {
			resultStr += string(result[i] + 'a')
		} else {
			resultStr += string(result[i] + '0')
		}
	}

	return resultStr
}