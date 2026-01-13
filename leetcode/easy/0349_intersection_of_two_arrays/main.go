package main

/*
# 349. Intersection of Two Arrays

Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.

Constraints:

	1 <= nums1.length, nums2.length <= 1000
	0 <= nums1[i], nums2[i] <= 1000
*/
func main() {
	// test := [8]byte{}
	// fmt.Printf("%08d", bitSet1[0])
}

// optimized intersection function

// https://leetcode.com/problems/intersection-of-two-arrays/description/
// /*
func intersection(nums1 []int, nums2 []int) []int {
	bitSet1 := [1000/8 + 1]byte{}
	result := make([]int, 0)

	for _, v := range nums1 {
		bIdx := v / 8
		bit := v % 8
		bitSet1[bIdx] |= 1 << bit
	}
	for _, v := range nums2 {
		bIdx := v / 8
		bit := v % 8
		if bitSet1[bIdx]&(1<<bit) != 0 {
			result = append(result, v)
			bitSet1[bIdx] ^= 1 << bit
		}
	}

	return result
}

// */

/* old intersection function (*benefit* return sorted array)
func intersection(nums1 []int, nums2 []int) []int {
	bitSet1 := [1000/8 + 1]byte{}
	bitSet2 := [1000/8 + 1]byte{}
	result := make([]int, 0)

	for _, v := range nums1 {
		bIdx := v / 8
		bit := v % 8
		bitSet1[bIdx] |= 1 << bit
	}
	for _, v := range nums2 {
		bIdx := v / 8
		bit := v % 8
		bitSet2[bIdx] |= 1 << bit

	}

	intersectionBits := [1000/8 + 1]byte{}
	for i, _ := range intersectionBits {
		intersectionBits[i] = bitSet1[i] & bitSet2[i]
	}

	for i, v := range intersectionBits {
		if v != 0 {
			for j := 0; j < 8; j++ {
				hasNum := v&(1<<j) != 0

				if hasNum {
					result = append(result, i*8+j)
				}
			}
		}
	}

	return result
}
*/
