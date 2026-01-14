package main

import (
	"fmt"
)

/*
# 27. Remove Element

Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

	Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
	Return k.
*/
func main() {
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}

	var valNums int = removeElement(nums1, 2)

	fmt.Printf("Array: %v\nCounter exeptional num: %d\n", nums1, valNums)

}

/*
removeElement removes all occurrences of val in nums in-place and returns the number of elements in nums which are not equal to val.

eg. removeElement([]int{3, 2, 2, 3}, 3) returns 2 and modifies the array to [2, 2, _, _].

link https://leetcode.com/problems/remove-element/description/
*/
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	right := len(nums) - 1
	left := 0

	for left <= right {
		if nums[left] == val && nums[right] != val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
			left++
			continue
		} else if nums[left] == val {
			right--
			continue
		}

		if nums[right] == val {
			right--
			left++
		} else {
			left++
		}
	}

	return left
}
