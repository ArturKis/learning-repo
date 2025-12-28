package main

import "fmt"

/*
# 350. Intersection of Two Arrays II

Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

Constraints:

  - 1 <= nums1.length, nums2.length <= 1000
  - 0 <= nums1[i], nums2[i] <= 1000

Follow up:

  - What if the given array is already sorted? How would you optimize your algorithm?
  - What if nums1's size is small compared to nums2's size? Which algorithm is better?
  - What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/
func main() {
	testCases := []struct {
		n1 []int
		n2 []int
		e  []int
	}{
		{
			n1: []int{1, 2, 2, 1},
			n2: []int{2, 2},
			e:  []int{2, 2},
		},
		{
			n1: []int{4, 9, 5},
			n2: []int{9, 4, 9, 8, 4},
			e:  []int{9, 4},
		},
		{
			n1: []int{1, 2, 3},
			n2: []int{4, 5, 6},
			e:  []int{},
		},
		{
			n1: []int{1},
			n2: []int{1},
			e:  []int{1},
		},
	}

	for i, tc := range testCases {
		res := intersect(tc.n1, tc.n2)
		fmt.Printf("Test %d:\n  Input 1: %v\n  Input 2: %v\n  Expected: %v\n  Result  : %v\n\n", i+1, tc.n1, tc.n2, tc.e, res)
	}
}

// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	arr := [1001]int{}
	result := []int{}

	for _, v := range nums1 {
		arr[v]++
	}

	for _, v := range nums2 {
		if arr[v] > 0 {
			result = append(result, v)
			arr[v]--
		}
	}

	return result
}

/*
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	m := make(map[int]int)
	result := []int{}

	for _, v := range nums1 {
		m[v]++
	}

	for _, v := range nums2 {
		if mv, ok := m[v]; ok && mv > 0 {
			result = append(result, v)
			m[v]--
		}
	}

	return result
}
*/
