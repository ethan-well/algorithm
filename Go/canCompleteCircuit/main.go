package main

import "log"

func main()  {
	gas := []int{2,3,4}
	cost := []int{3,4,3}
	log.Printf("i: %d", canCompleteCircuit(gas, cost))
}


func canCompleteCircuit(gas []int, cost []int) int {

	for i := 0; i < len(gas); i ++ {
		if canComplete(i, gas, cost) {
			return i
		}
	}

	return -1
}

func canComplete(n int, gas, cost []int) bool {
	gasHas := 0

	for i, g := range gas[n:] {
		gasHas = gasHas + g - cost[n:][i]
		if gasHas < 0 {
			return false
		}
	}

	for i, g := range gas[:n] {
		gasHas = gasHas + g - cost[:n][i]
		if gasHas < 0 {
			return false
		}
	}

	return true
}