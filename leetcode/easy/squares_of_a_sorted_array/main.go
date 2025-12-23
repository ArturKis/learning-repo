package main

import "fmt"

/*
# Squares of a Sorted Array

Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

# Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?
*/
func main() {
	// nums1 := []int{-4, -1, 0, 3, 10}
	// fmt.Println(sortedSquares(nums1))

	nums2 := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums2))
}

// https://leetcode.com/problems/squares-of-a-sorted-array/submissions/1851645108/
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))

	left := 0
	right := len(nums) - 1

	for i := right; left <= right; i-- {
		squareL, squareR := nums[left]*nums[left], nums[right]*nums[right]

		if squareL > squareR {
			result[i] = squareL
			left++
		} else {
			result[i] = squareR
			right--
		}

	}

	return result
}
