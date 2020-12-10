package main

func main() {
	println(findKthNumber(3, 2))
}

func findKthNumber(n int, k int) int {
	//prefix := currentNode(1,  100)

	count := 1
	prefix := 1
	for count < k {
		c := currentNode(prefix, n)
		if count + c > k {
			count += 1
			prefix *= 10
		} else {
			count += c
			prefix += 1
		}
	}

	return prefix
}

func currentNode(prefixN int, n int) int {
	nextPrefixN := prefixN + 1
	var count int
	for prefixN <= n  {
		if nextPrefixN > n {
			nextPrefixN = n+1
		}

		count += nextPrefixN - prefixN

		prefixN *= 10
		nextPrefixN *= 10
	}

	return count
}