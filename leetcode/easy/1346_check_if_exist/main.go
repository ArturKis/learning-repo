package main

import "fmt"

/*
# 1346. Check If N and Its Double Exist

Given an array arr of integers, check if there exist two indices i and j such that :
  - i != j
  - 0 <= i, j < arr.length
  - arr[i] == 2 * arr[j]
*/
func main() {
	arr1 := []int{10, 2, 5, 3}
	result1 := checkIfExist(arr1)
	fmt.Println(result1)

	arr2 := []int{3, 1, 7, 11}
	result2 := checkIfExist(arr2)
	fmt.Println(result2)
}

// https://leetcode.com/problems/check-if-n-and-its-double-exist/description/
func checkIfExist(arr []int) bool {
	if len(arr) < 2 || arr == nil {
		return false
	}

	for i, v := range arr {
		for j := i + 1; j < len(arr); j++ {
			if v != arr[j] && v == 2*arr[j] || v*2 == arr[j] {
				return true
			}
		}
	}

	return false
}
