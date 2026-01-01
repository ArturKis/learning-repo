package main

import "fmt"

/*
# 454. 4Sum II

Given four integer arrays nums1, nums2, nums3, and nums4 all of length n, return the number of tuples (i, j, k, l) such that:

  - 0 <= i, j, k, l < n
  - nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

Constraints:

  - n == nums1.length
  - n == nums2.length
  - n == nums3.length
  - n == nums4.length
  - 1 <= n <= 200
  - -2^28 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2^28
*/
func main() {
	testCases := []struct {
		n1 []int
		n2 []int
		n3 []int
		n4 []int
		e  int
	}{
		{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
		{[]int{0}, []int{0}, []int{0}, []int{0}, 1},
	}

	for i, tc := range testCases {
		result := fourSumCount(tc.n1, tc.n2, tc.n3, tc.n4)
		fmt.Printf("Test #%d\nExpected %d, got %d\n\n", i+1, tc.e, result)
	}
}

// https://leetcode.com/problems/4sum-ii/description/
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int, len(nums1)*len(nums1))
	result := 0

	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			m[v1+v2]++
		}
	}

	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			mk := -(v3 + v4)
			if c, ok := m[mk]; ok {
				result += c
			}
		}
	}

	return result
}
