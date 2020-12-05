package main

import "log"

func main() {
	if isHappy(7) {
		log.Printf("happy number")
	} else {
		log.Printf("not happy")
	}
}


func isHappy(n int) bool {
	var resultArr []int
	for  {
		n = getSum(n)
		if n == 1 {
			return true
		}
		if nInResultArr(n, resultArr) {
			return false
		}
		resultArr = append(resultArr, n)
	}
}

func getSum(n int) int {
	var sum int
	for n != 0 {
		t := n % 10
		n = n / 10
		sum += t * t
	}

	return sum
}

func nInResultArr(n int, resultArr[]int) bool {
	for _, v := range resultArr {
		if v == n {
			return true
		}
	}

	return false
}