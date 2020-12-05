package main

import "log"

func main() {
	log.Printf("sum: %d", getSum(1111, 333333))
}

func getSum(a int, b int) int {
	sum1 := a ^ b
	sum2 := a & b << 1
	if sum2 == 0 {
		return sum1
	}
	return getSum(sum1, sum2)
}
