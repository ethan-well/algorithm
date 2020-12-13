package main

func main() {
	// subStr := longestPalindrome("b")
	// subStr := longestPalindrome("babad")

	subStr := longestPalindrome("aaaa")
	println(subStr)
}

func longestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}

	var start, end int
	for i := 0; i < len(s); i ++ {
		l1, r1 := getLRIndexOfCurrentChar(s, i, i)
		l2, r2 := getLRIndexOfCurrentChar(s, i, i + 1)
		if r1 - l1 > end - start {
			start, end = l1, r1
		}

		if r2 - l2 > end - start {
			start, end = l2, r2
		}
	}

	return s[start: end+1]
}

func getLRIndexOfCurrentChar(s string, l, r int) (int, int) {

	for ; l >= 0 && r < len(s) && s[l] == s[r]; {
		l --
		r ++
	}

	return l + 1, r -1
}