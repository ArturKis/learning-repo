package main

import (
	"fmt"
)

/*
# 219. Contains Duplicate II

Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

Constraints:

  - 1 <= nums.length <= 10^5
  - -10^9 <= nums[i] <= 10^9
  - 0 <= k <= 10^5
*/
func main() {
	testCases := []struct {
		n []int
		k int
		e bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for i, tc := range testCases {
		result := containsNearbyDuplicate(tc.n, tc.k)
		fmt.Printf("Test #%d (expected: %t) => %t\nTarget distance: %d\nInput:  nums:%v\n\n", i+1, tc.e, result, tc.k, tc.n)
	}
}

// https://leetcode.com/problems/contains-duplicate-ii/description/
func containsNearbyDuplicate(nums []int, k int) bool {
	c := min(len(nums), k+1)
	m := make(map[int]struct{}, c)

	for i, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}

		m[v] = struct{}{}

		if len(m) > k {
			delete(m, nums[i-k])
		}
	}

	return false
}

/*
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))

	for i, v := range nums {
		if mv, ok := m[v]; ok && i-mv <= k {
			return true
			} else {
			m[v] = i
		}
	}

	return false
}
*/
