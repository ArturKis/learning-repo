package main

import (
	"fmt"
)

/*
# Find All Numbers Disappeared in an Array

Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

# Follow up: Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
*/
func main() {
	nums1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result1 := findDisappearedNumbers(nums1)
	fmt.Println(result1)
}

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/
func findDisappearedNumbers(nums []int) []int {

	for _, v := range nums {
		var target int

		if v < 0 {
			target = -v - 1
		} else {
			target = v - 1
		}

		if nums[target] > 0 {
			nums[target] = -nums[target]
		}
	}

	result := make([]int, 0)

	for i, v := range nums {
		if v > 0 {
			result = append(result, i+1)
		}
	}

	return result
}
