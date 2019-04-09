package main

import "fmt"

func areaArountCurentPoint(y, x int, grid [][]int) int {
	area := 0
	maxY := len(grid)
	maxX := len(grid[1])

	if v := grid[y][x]; v == 0 {
		return area
	}

	grid[y][x] = -1

	area = 1

	if xL := x - 1; 0 <= xL && grid[y][xL] != -1 {
		area += areaArountCurentPoint(y, xL, grid)
	}

	if xR := x + 1; xR <= maxX-1 && grid[y][xR] != -1 {
		area += areaArountCurentPoint(y, xR, grid)
	}

	if yT := y - 1; 0 <= yT && grid[yT][x] != -1 {
		area += areaArountCurentPoint(yT, x, grid)
	}

	if yB := y + 1; yB < maxY && grid[yB][x] != -1 {
		area += areaArountCurentPoint(yB, x, grid)
	}

	return area
}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for y, arr := range grid {
		for x, _ := range arr {
			if result := areaArountCurentPoint(y, x, grid); result > maxArea {
				maxArea = result
			}
		}
	}

	return maxArea
}

func main() {
	var arr = [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(arr))
}
