package main

import (
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"
)

/*
# 119. Pascal's Triangle II

Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown: https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif

Constraints:
  - 0 <= rowIndex <= 33

Follow up: Could you optimize your algorithm to use only O(rowIndex) extra space?
*/
func main() {
	testCases := []struct {
		input  int
		output []int
	}{
		{input: 0, output: []int{1}},
		{input: 1, output: []int{1, 1}},
		{input: 3, output: []int{1, 3, 3, 1}},
		{input: 5, output: []int{1, 5, 10, 10, 5, 1}},
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	// defer tw.Flush()

	for i, tc := range testCases {
		result := getRow(tc.input)
		status := "PASS"

		if !reflect.DeepEqual(tc.output, result) {
			status = "FAIL"
		}

		fmt.Fprintf(tw, "Test: %d\n", i+1)
		fmt.Fprintf(tw, "======================\n")
		fmt.Fprintf(tw, "Input(triangle index):\t%d\n", tc.input)
		fmt.Fprintf(tw, "Expected:\t%v\n", tc.output)
		fmt.Fprintf(tw, "Result:\t%v\n", result)
		fmt.Fprintf(tw, "Test:\t%s\n\n", status)
	}

}

// recursion solution
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	result := append([]int{1}, getRow(rowIndex-1)...)
	for i := 1; i < rowIndex; i++ {
		result[i] += result[i+1]
	}

	return result
}

/*
// after refactor
// https://leetcode.com/problems/pascals-triangle-ii/description/
func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		result[i] = 1
	}

	for i := 1; i < rowIndex; i++ {
		for j := 0; j < rowIndex-i; j++ {
			result[j+1] += result[j]
		}
	}

	return result
}
*/

/*
// it took 2 hours on an easy task :(
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	result := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		result[i] = 1
	}

	for i := 1; rowIndex-i > 0; i++ {
		for j := 0; rowIndex-i-j > 0; j++ {
			result[j+1] = result[j] + result[j+1]
		}
	}

	return result
}
*/

/*
index=6 n=index+1
[1,1, 1, 1, 1,1,1]
[1,2, 3, 4, 5,6,1]  index-1
[1,3, 6,10,15,6,1]  index-2
[1,4,10,20,15,6,1]  index-3
[1,5,15,20,15,6,1]  index-4
[1,6,15,20,15,6,1]  index-5
*/
