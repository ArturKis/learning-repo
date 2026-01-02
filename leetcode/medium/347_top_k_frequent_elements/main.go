package main

import (
	"fmt"
)

/*
# 347. Top K Frequent Elements

Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Constraints:

  - 1 <= nums.length <= 10^5
  - -10^4 <= nums[i] <= 10^4
  - k is in the range [1, the number of unique elements in the array].
  - It is guaranteed that the answer is unique.

Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/
func main() {
	testCases := []struct {
		a []int
		k int
		e []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
	}

	for i, tc := range testCases {
		result := topKFrequent(tc.a, tc.k)
		fmt.Printf("Test: %d\ntopKFrequent(%v, %d) = %v (%v)\n\n", i+1, tc.a, tc.k, result, tc.e)

	}
}

// https://leetcode.com/problems/top-k-frequent-elements/description/
func topKFrequent(nums []int, k int) []int {
	mc := make(map[int]int, len(nums))
	result := make([]int, 0, k)

	for _, v := range nums {
		mc[v]++
	}

	slice := make([][]int, len(nums)+1)

	for k, v := range mc {
		slice[v] = append(slice[v], k)
	}

	for i := len(slice) - 1; i >= 0 && len(result) < k; i-- {
		if len(slice[i]) > 0 {
			result = append(result, slice[i]...)
		}
	}

	return result[:k]
}
