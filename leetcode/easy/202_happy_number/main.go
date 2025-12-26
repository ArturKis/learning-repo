package main

import (
	"fmt"
	"math"
)

/*
# Happy Number

Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

  - Starting with any positive integer, replace the number by the sum of the squares of its digits.
  - Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
  - Those numbers for which this process ends in 1 are happy.

Return true if n is a happy number, and false if not.

Constraints:

  - 1 <= n <= 2^31 - 1
*/
func main() {
	nums := []int{19, 2, 223123, int(math.Pow(2, 31)) - 1}

	printResults(nums)
}

// https://leetcode.com/problems/happy-number/description/
func isHappy(n int) bool {
	bitSet := [128]byte{}

	for n != 1 {
		sum := 0

		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}

		n = sum

		bIdx := n / 8
		bit := n % 8

		if (bitSet[bIdx] & (1 << bit)) != 0 {
			return false
		}

		bitSet[bIdx] |= 1 << bit
	}

	return true
}

func printResults(nums []int) {
	for _, v := range nums {
		fmt.Printf("A number %d is happy? - %t\n", v, isHappy(v))
	}
}
