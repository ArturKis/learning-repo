package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

/*
# 509. Fibonacci Number

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

Constraints:
  - 0 <= n <= 30
*/
func main() {
	// /*
	testCases := []struct {
		input    int
		expected int
	}{
		{input: 2, expected: 1},
		{input: 3, expected: 2},
		{input: 4, expected: 3},
		// {input: 4, expected: 0}, // fail test
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)
	defer tw.Flush()
	fmt.Fprintf(tw, "ID\tINPUT\tRESULT\tEXPECTED\tSTATUS\n")

	for i, tc := range testCases {
		result := fib(tc.input)
		status := "PASS"

		if result != tc.expected {
			status = "FAIL"
		}

		fmt.Fprintf(tw, "Test: %d\t%d\t%d\t%d\t%s\n", i+1, tc.input, result, tc.expected, status)
	}
	// */

	/*
		testCases := []struct {
			num int
			pow int
		}{
			{5, 10}, // 9765625
			{2, 8},  // 256
		}

		tw := tabwriter.NewWriter(os.Stdout, 0, 1, 0, ' ', tabwriter.Debug)
		defer tw.Flush()
		fmt.Fprintf(tw, "\tNUMBER\tPOWER\tRESULT\n")

		for i, tc := range testCases {
			result := sketch(tc.num, tc.pow)
			fmt.Fprintf(tw, "Test: %d\t%d\t%d\t%d\n", i+1, tc.num, tc.pow, result)
		}
	*/
}

// https://leetcode.com/problems/fibonacci-number/description/
func fib(n int) int {
	if n < 2 {
		return n
	}

	result := [4]int{1, 0, 0, 1}
	base := [4]int{1, 1, 1, 0}

	var fn func(n int, r [4]int, b [4]int) [4]int
	fn = func(n int, result, base [4]int) [4]int {
		if n == 0 {
			return result
		}

		if n%2 != 0 {
			result = multiplyMatrices2x2(result, base)
		}

		base = multiplyMatrices2x2(base, base)
		n /= 2

		return fn(n, result, base)
	}

	//  |n+1   n|	[0, 1,
	//  |n   n-1|	 2, 3]
	return fn(n, result, base)[1]
}

func multiplyMatrices2x2(a, b [4]int) [4]int {
	var result [4]int

	result[0] = a[0]*b[0] + a[1]*b[2]
	result[1] = a[0]*b[1] + a[1]*b[3]

	result[2] = a[2]*b[0] + a[3]*b[2]
	result[3] = a[2]*b[1] + a[3]*b[3]

	return result
}

/* ... from my 99_random_practice/power_impl(recurcion_practice) dir
func sketch(num int, power int) int {
	var fn func(b int, p int, r int) int

	fn = func(base, p, result int) int {
		if p == 0 {
			return result
		}

		if p%2 != 0 {
			result = result * base
		}

		base *= base
		p /= 2

		return fn(base, p, result)
	}

	return fn(num, power, 1)
}
*/

/*
func fib(n int) int {
	if n < 2 {
		return n
	}

	cache := make([]int, n+1)
	var fn func(int) int

	fn = func(n int) int {
		if n < 2 {
			return n
		}

		if v := cache[n]; v != 0 {
			return v
		}

		result := fn(n-1) + fn(n-2)
		cache[n] = result

		return result
	}

	return fn(n)
}
*/
