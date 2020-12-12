package main

import "strings"

func main() {

	s := "abcabcbb"
	println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	var start, end, count int

	for i := 0; i < len(s); i ++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			end ++
		} else {
			start += index + 1
		}

		count = max(count, end - start)
	}

	return count
}

func max(n, m int) int {
	if n >= m {
		return n
	}

	return m
}