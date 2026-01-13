package main

import "fmt"

/*
# 1299. Replace Elements with Greatest Element on Right Side

Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.

After doing so, return the array.
*/
func main() {
	slice1 := []int{17, 18, 5, 4, 6, 1}
	result1 := replaceElements(slice1)
	fmt.Println("Result1:", result1)

	slice2 := []int{12}
	result2 := replaceElements(slice2)
	fmt.Println("Result2:", result2)
}

// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/description/
func replaceElements(arr []int) []int {
	if len(arr) > 0 {
		lastIdx := len(arr) - 1
		max := arr[lastIdx]

		for i := lastIdx - 1; i >= 0; i-- {
			if max < arr[i] {
				max, arr[i] = arr[i], max
			} else {
				arr[i] = max
			}
		}

		arr[lastIdx] = -1
	}

	return arr
}
