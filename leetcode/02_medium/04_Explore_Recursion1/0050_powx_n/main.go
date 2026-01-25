package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

/*
# 50. Pow(x, n)

Implement pow(x, n), which calculates x raised to the power n (i.e., x^n).

Constraints:
  - -100.0 < x < 100.0
  - -2^31 <= n <= 2^31-1
  - n is an integer.
  - Either x is not zero or n > 0.
  - -10^4 <= x^n <= 10^4
*/
func main() {
	testCases := []struct {
		x float64
		n int
		e float64
	}{
		{x: 2.00000, n: 10, e: 1024.00000},
		{x: 2.10000, n: 3, e: 9.26100},
		{x: 2.00000, n: -2, e: 0.25000},
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	defer tw.Flush()

	fmt.Fprintln(tw, "ID\tNUMBER\tPOWER\tEXPECTED\tRESULT\tSTATUS")

	for i, tc := range testCases {
		result := myPow(tc.x, tc.n)
		status := "PASS"

		diff := result - tc.e

		if max(diff, -diff) > 1e-9 {
			status = "FAIL"
		}

		fmt.Fprintf(tw, "Test: %d\t%.5f\t%d\t%.5f\t%.5f\t%s\n", i+1, tc.x, tc.n, tc.e, result, status)
	}
}

// https://leetcode.com/problems/powx-n/description/
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	p := n
	if p < 0 {
		x = 1 / x
		p = -p
	}

	var fn func(float64, float64, int) float64
	fn = func(acc, base float64, pow int) float64 {
		if pow == 0 {
			return acc
		}

		if pow%2 != 0 {
			acc *= base
		}

		return fn(acc, base*base, pow/2)
	}

	return fn(1.0, x, p)
}
