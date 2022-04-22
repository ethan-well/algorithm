package oom

import "strings"

func main() {
	var str0 = "1212121212121212"
	// 内存泄漏
	str1 := str0[:10]

	str2 := (" " + str0[:10])[:10]

	str3 := string([]byte(str0[:10]))

	var builder strings.Builder
	builder.Grow(10)
	builder.WriteString(str0[:10])
	str4 := builder.String()

	str5 := strings.Repeat(str0[:10], 1)


	s0 := []int{1,2,3,4,5,6,7,8,9}
	// 内存泄漏
	s1 := s0[:5]
	// 解决方案
	s1 := append(s0[:0], s0[:5]...)
}
