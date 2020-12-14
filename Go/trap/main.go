package main

func main() {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	println(trap(height))

	height2 := []int{4,2,0,3,2,5}
	println(trap(height2))
}

func trap(height []int) int {
	var water int
	for i := 0; i <= len(height); i ++ {
		waterI := getWater(i, height)
		water += waterI
	}

	return water
}

func getWater(i int, height []int) int {
	if i == 0 || i == len(height) {
		return 0
	}

	var lMax, rMax int

	for _, h := range height[:i] {
		if h > lMax {
			lMax = h
		}
	}

	for _, h := range height[i+1:] {
		if h > rMax {
			rMax = h
		}
	}

	var water int
	if lMax > rMax {
		water = rMax - height[i]
	} else {
		water = lMax - height[i]
	}

	if water < 0 {
		return 0
	}

	return water
}


