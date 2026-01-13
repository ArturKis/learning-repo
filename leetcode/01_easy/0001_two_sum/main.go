package main

import "fmt"

/*
# Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Constraints:

  - 2 <= nums.length <= 10^4
  - -10^9 <= nums[i] <= 10^9
  - -10^9 <= target <= 10^9
  - Only one valid answer exists.

Follow-up: Can you come up with an algorithm that is less than O(n^2) time complexity?
*/
func main() {
	testCases := []struct {
		n []int
		t int
		e []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range testCases {
		result := twoSum(tc.n, tc.t)
		fmt.Printf("Array: %v | Target num: %d | Expected: %v | Result: %v\n", tc.n, tc.t, tc.e, result)
	}
}

// https://leetcode.com/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		if mv, ok := m[target-v]; ok {
			return []int{mv, i}
		}
		m[v] = i
	}

	return nil
}
