package main

import "fmt"

/*
# Third Maximum Number

Given an integer array nums, return the third distinct maximum number in this array. If the third maximum does not exist, return the maximum numbe
*/
func main() {

	// /*
	nums1 := []int{3, 2, 10, 5, 4, 10, 5} // (4)
	fmt.Println(thirdMax(nums1))

	nums2 := []int{2, 2, 1} // (2)
	fmt.Println(thirdMax(nums2))

	nums3 := []int{-1, -2, -3}
	fmt.Println(thirdMax(nums3))

	nums4 := []int{10, 10, 10}
	fmt.Println(thirdMax(nums4))
	// */

	nums5 := []int{1, -2147483648, 2}
	fmt.Println(thirdMax(nums5))
}

// https://leetcode.com/problems/third-maximum-number/description/
func thirdMax(nums []int) int {
	var max1, max2, max3 int
	var hasM1, hasM2, hasM3 bool

	for _, v := range nums {
		if !hasM1 {
			max1 = v
			hasM1 = true
		} else if max1 < v {
			max1, max2, max3 = v, max1, max2
			if hasM2 {
				hasM3 = true
			}
			hasM2 = true
			continue
		}

		if !hasM2 && v != max1 {
			max2 = v
			hasM2 = true
		} else if hasM2 && max2 < v && v != max1 {
			max2, max3 = v, max2
			hasM3 = true
			continue
		}

		if !hasM3 && v != max1 && v != max2 {
			max3 = v
			hasM3 = true
		} else if hasM3 && max3 < v && v != max1 && v != max2 {
			max3 = v
		}

	}

	if hasM3 {
		return max3
	}

	return max1
}
