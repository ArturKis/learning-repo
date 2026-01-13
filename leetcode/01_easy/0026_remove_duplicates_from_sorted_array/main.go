package main

import "fmt"

/*
# 26. Remove Duplicates from Sorted Array

Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same.

Consider the number of unique elements in nums to be k. After removing duplicates, return the number of unique elements k.

The first k elements of nums should contain the unique numbers in sorted order. The remaining elements beyond index k - 1 can be ignored.
*/
func main() {
	slice := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	k := removeDuplicates(slice)

	fmt.Printf("Array: %v\nLength without duplicates: %d\n", slice, k)
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	target := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[target] = nums[i]
			target++
		}
	}

	return target
}
