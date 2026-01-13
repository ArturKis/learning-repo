package main

import "fmt"

/*
# 941. Valid Mountain Array

Given an array of integers arr, return true if and only if it is a valid mountain array.

Recall that arr is a mountain array if and only if:

	arr.length >= 3
	There exists some i with 0 < i < arr.length - 1 such that:
	   	arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
	    arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
*/
func main() {
	// /*
	slice1 := []int{1, 3, 5, 7, 6, 4, 2} // true
	result1 := validMountainArray(slice1)
	fmt.Printf("Result slice1 is: %v\n", result1)

	slice2 := []int{2, 4, 6, 8, 10} // false
	result2 := validMountainArray(slice2)
	fmt.Printf("Result slice2 is: %v\n", result2)
	// */

	slice3 := []int{1, 3, 2} // true
	result3 := validMountainArray(slice3)
	fmt.Printf("Result slice3 is: %v\n", result3)
}

// https://leetcode.com/problems/valid-mountain-array/description/
func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	isInc := true
	prevValue := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] == prevValue || (arr[i] > prevValue && !isInc) || (i == len(arr)-1 && isInc) || (i == 1 && arr[i] < prevValue) {
			return false
		}

		if isInc && arr[i+1] < arr[i] {
			isInc = false
		}

		prevValue = arr[i]

	}

	return true
}
