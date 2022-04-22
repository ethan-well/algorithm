package myAtoi

import (
	"errors"
	"strings"
)

func AToi(input string) (int64, error) {
	err := errors.New("Input is not a int64 number")

	input = strings.TrimSpace(input)
	inputLen := len(input)
	if inputLen == 0 {
		return 0, err
	}

	// base 用来标记正负数
	var base = 1
	if c := input[0]; c == '-' || c == '+' {
		// 标记为负数
		if c == '-' {
			base = -1
		}

		input = input[1:]
		inputLen = inputLen - 1
	}

	// int64 最大，最小数的长度为 19
	maxLen := 19
	// 输入字符串长度溢出判断
	if inputLen < 1 || inputLen > maxLen {
		return 0, err
	}

	var number int64
	for i, char := range []byte(input) {
		n := char - '0'
		if n > 9 || n < 0 {
			return 0, err
		}

		// 判断正数是否超过 max int64
		// 只对输入字符串长度等于 maxLen: 19 的进行检查
		// 长度超过 19 的在外层已经判断
		if inputLen == maxLen && base == 1 && overflowMax(char, i) {
			return 0, errors.New("Input is overflow max Int64")
		}

		// 判断负数是否小于 min int64
		// 只对输入字符串长度等于 maxLen: 19 的进行检查
		// 长度超过 19 的在外层已经判断
		if inputLen == maxLen && base == -1 && overflowMin(char, i) {
			return 0, errors.New("input is overflow min Int64")
		}

		//number += int64(n) * int64(math.Pow10(inputLen-i-1))
		number = number*10 + int64(n)
	}

	return number * int64(base), nil
}

// 判断正数是否超过 max int64
func overflowMax(char byte, index int) bool {
	maxIn64Str := []byte("9223372036854775807")

	return maxIn64Str[index]-char > 0
}

// 判断负数是否超过 min int64
func overflowMin(char byte, index int) bool {
	minIn64Str := []byte("9223372036854775808")

	return minIn64Str[index]-char > 0
}
