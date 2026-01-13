package main

import (
	"fmt"
	"slices"
)

/*
# 1051. Height Checker

A school is trying to take an annual photo of all the students. The students are asked to stand in a single file line in non-decreasing order by height. Let this ordering be represented by the integer array expected where expected[i] is the expected height of the ith student in line.

You are given an integer array heights representing the current order that the students are standing in. Each heights[i] is the height of the ith student in line (0-indexed).

Return the number of indices where heights[i] != expected[i].
*/
func main() {
	heights1 := []int{1, 1, 4, 2, 1, 3} // [1, 1, 1, 2, 3, 4,] (3)
	fmt.Println(heightChecker(heights1))
}

// https://leetcode.com/problems/height-checker/description/
func heightChecker(heights []int) int {
	if len(heights) <= 1 {
		return 0
	}

	expected := slices.Clone(heights)

	for i, _ := range expected {
		for j := i + 1; j < len(expected); j++ {
			if expected[i] > expected[j] {
				expected[i], expected[j] = expected[j], expected[i]
			}
		}
	}

	notMatch := 0

	for i, v := range heights {
		if v != expected[i] {
			notMatch++
		}
	}

	return notMatch
}
