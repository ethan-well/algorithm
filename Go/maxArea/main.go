package main

func main() {
	height := []int{1,8,6,2,5,4,8,3,7}
	println(maxAreaV2(height))
}

func maxArea(height []int) int {
	maxArea := 0
	for i := 0; i < len(height); i ++ {
		y := height[i]
		if i == 1 {
			println("yi", y)
		}

		for j := i +1; j < len(height); j ++ {
			x := j - i

			if i == 1 && j == 8 {
				println(y, i, x)
			}

			if y > height[j] {
				y = height[j]
			}

			area := x * y
			if area > maxArea {
				maxArea = area
			}

			y = height[i]
		}
	}

	return maxArea
}


func maxAreaV2(height []int) int {
	i, j, maxArea := 0, len(height) -1, 0

	for i < j {
		area := 0
		if height[i] < height[j] {
			area = height[i] * (j - i)
			i ++
		} else {
			area = height[j] * (j - i)
			j --
		}

		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
