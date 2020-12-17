package main

import "fmt"

import "strconv"

func main()  {
	str1 := "98123"
	str2 := "12345"

	println(addStrings(str1,str2))
}

func addStrings(num1 string, num2 string) string {

	var nextN int
	var result string
	for i := 0; i < len(num1) || i < len(num2); i ++ {
		n1, n2 := 0, 0

		if i < len(num1) {
			n1, _ = strconv.Atoi(string(num1[len(num1) - i -1]))
		}

		if i < len(num2) {
			n2, _ = strconv.Atoi(string(num2[len(num2) -i - 1]))
		}

		sum := n1 + n2 + nextN

		result =  fmt.Sprintf("%d", sum % 10) + result
		nextN = sum / 10
	}

	if nextN > 0 {
		result =  fmt.Sprintf("%d", nextN) + result
	}

	return result
}
