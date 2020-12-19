package main

func main() {

	println(isValid("()[]{}"))
}

func isValid(s string) bool {
	var parentheses []byte
	for i := 0; i < len(s); i ++ {

		if i == 0 {
			parentheses = append(parentheses, s[i])
			continue
		}

		if s[i] == '(' ||
			s[i] == '{' ||
			s[i] == '[' {

			parentheses = append(parentheses, s[i])
			continue
		}

		var expected byte
		if s[i] == ')' {
			expected = parentheses[len(parentheses)-1] + 1
		} else {
			expected = parentheses[len(parentheses)-1] + 2
		}

		if expected != s[i] {
			return false
		}

		parentheses = parentheses[:len(parentheses)-1]
	}

	if len(parentheses) != 0 {
		return false
	}

	return true
}
