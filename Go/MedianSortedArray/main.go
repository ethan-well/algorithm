package main

func main()  {
	nums1 := []int{2,3,4,5,6,7,8,9,10}
	nums2 := []int{1}
	println(findMedianSortedArrays(nums1, nums2))

	println(findKNodeFromArrays(nums1, nums2, 8))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return getMedianFromArray(nums2)
	}

	if len(nums2) == 0 {
		return getMedianFromArray(nums1)
	}

	if len(nums1) + len(nums2) == 2 {
		return float64(nums1[0] + nums2[0]) / 2.0
	}

	length := len(nums1) + len(nums2)
	if  length % 2 == 0 {
		k1 := length / 2
		m1 := findKNodeFromArrays(nums1, nums2, k1)

		k2 := length / 2 + 1
		m2 := findKNodeFromArrays(nums1, nums2, k2)

		return (m1 + m2) / 2
	} else {
		k := length / 2 + 1
		return findKNodeFromArrays(nums1, nums2, k)
	}
}

func findKNodeFromArrays(nums1, nums2 []int, k int)  float64 {
	if k < 0 {
		return 0
	}

	if len(nums1) + len(nums2) < k {
		return 0
	}

	if len(nums1) == 0 {
		return float64(nums2[k-1])
	}

	if len(nums2) == 0 {
		return float64(nums1[k-1])
	}

	if k == 1 {
		if nums1[0] < nums2[0] {
			return float64(nums1[0])
		}
		return float64(nums2[0])
	}

	t := k / 2

	i, j := t-1, t-1
	if len(nums1) < t {
		i = len(nums1) - 1
		j = t - len(nums1) - 1
	}

	if len(nums2) < t {
		j = len(nums2) - 1
		i = t - len(nums2) - 1
	}

	if nums1[i] <= nums2[j] {
		k = k - len(nums1[0:i+1])
		nums1 = nums1[i+1:]
	} else {
		k = k - len(nums2[0:j+1])
		nums2 = nums2[j+1:]
	}

	return findKNodeFromArrays(nums1, nums2, k)
}

func getMedianFromArray(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}

	if l := len(nums); l % 2 == 0 {
		return float64(nums[l/2] + nums[l/2-1]) / 2
	} else {
		return float64(nums[l/2])
	}
}