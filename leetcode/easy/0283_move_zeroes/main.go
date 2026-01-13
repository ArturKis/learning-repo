package main

import "fmt"

/*
# 283. Move Zeroes

Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.
*/
func main() {
	slice1 := []int{0, 0, 1, 2, 3, 0, 0, 3, 12}
	moveZeroes(slice1)
	fmt.Println(slice1)
}

// https://leetcode.com/problems/move-zeroes/
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	var target int
	findZero := false

	for i, v := range nums {
		if v != 0 && findZero {
			nums[i], nums[target] = nums[target], nums[i]
			target++
			continue
		}
		if !findZero && v == 0 {
			target = i
			findZero = true
		}
	}
}
