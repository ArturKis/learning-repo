package main

import "fmt"

/*
# 217. Contains Duplicate

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Constraints:

	1 <= nums.length <= 10^5
	-10^9 <= nums[i] <= 10^9
*/
func main() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 3, 4}
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Printf("Slice nums1 contains duplicates: %t\n", containsDuplicate(nums1))
	fmt.Printf("Slice nums2 contains duplicates: %t\n", containsDuplicate(nums2))
	fmt.Printf("Slice nums3 contains duplicates: %t\n", containsDuplicate(nums3))
}

// https://leetcode.com/problems/contains-duplicate/description/
func containsDuplicate(nums []int) bool {
	new := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := new[v]; ok {
			return true
		}

		new[v] = struct{}{}
	}

	return false
}
