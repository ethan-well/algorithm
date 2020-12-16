package main

func main()  {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	var sum int
	max := nums[0]
	for i := 0; i < len(nums); i ++ {
		if sum + nums[i] > sum {
			nums[i] += sum

			if nums[i] > max {
				max = nums[i]
			}
		}
		sum = nums[i]
	}

	return max
}
