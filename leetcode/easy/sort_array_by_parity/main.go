package main

import "fmt"

/*
# Sort Array By Parity

Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.

# Return any array that satisfies this condition

Explanation: The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
*/
func main() {
	slice1 := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(slice1))

	slice2 := []int{1, 0, 3}
	fmt.Println(sortArrayByParity(slice2))
}

// https://leetcode.com/problems/sort-array-by-parity/description/
func sortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left]%2 > nums[right]%2 {
			nums[left], nums[right] = nums[right], nums[left]
			right--
			left++
		}

		if nums[left]%2 == 0 {
			left++
		}

		if nums[right]%2 == 1 {
			right--
		}
	}

	return nums
}
