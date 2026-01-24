package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

/*
# 70. Climbing Stairs

You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Constraints:
  - 1 <= n <= 45
*/
func main() {
	testCases := []struct {
		input    int
		expected int
	}{
		{input: 2, expected: 2},
		{input: 3, expected: 3},
		{input: 4, expected: 5},
		{input: 5, expected: 8},
	}

	tw := tabwriter.NewWriter(os.Stdout, 10, 0, 1, ' ', tabwriter.Debug)
	defer tw.Flush()

	fmt.Fprintf(tw, "ID\tINPUT\tEXPECTED\tRESULT\tSTATUS\n")

	for i, tc := range testCases {
		result := climbStairs(tc.input)
		status := "PASS"

		if tc.expected != result {
			status = "FAIL"
		}

		fmt.Fprintf(tw, "Test: %d\t%d\t%d\t%d\t%s\n", i+1, tc.input, tc.expected, result, status)
	}
}

// /* Main Solution
// https://leetcode.com/problems/climbing-stairs/description/
func climbStairs(n int) int {
	cache := make([]int, n+1)

	var fn func(int) int
	fn = func(n int) int {
		if n <= 2 {
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

// */

/*
n=5

n0
.

n1
1. 1
1way

n2
1. 1-1
2. 2
2ways

n3
step 1: possible 2 | 3-2=1 | n1 include all (1) + 2
step 2: n2 include all(2 || 1 + 1) + 1

n1+n2
...
*/

/* balovstvo na go
// ================================================================================

type MyMatrixInt int

func (a MyMatrixInt) Add(b MyMatrixInt) MyMatrixInt { return a + b }
func (a MyMatrixInt) Mul(b MyMatrixInt) MyMatrixInt { return a * b }

// ================================================================================

type MyMatrixUint64 uint64

func (a MyMatrixUint64) Add(b MyMatrixUint64) MyMatrixUint64 { return a + b }
func (a MyMatrixUint64) Mul(b MyMatrixUint64) MyMatrixUint64 { return a * b }

// ================================================================================

type MyMatrixFloat64 float64

func (a MyMatrixFloat64) Add(b MyMatrixFloat64) MyMatrixFloat64 { return a + b }
func (a MyMatrixFloat64) Mul(b MyMatrixFloat64) MyMatrixFloat64 { return a * b }

// ================================================================================

type MyMatrixBigInt struct{ *big.Int }

func (a MyMatrixBigInt) Add(b MyMatrixBigInt) MyMatrixBigInt {
	return MyMatrixBigInt{new(big.Int).Add(a.Int, b.Int)}
}

func (a MyMatrixBigInt) Mul(b MyMatrixBigInt) MyMatrixBigInt {
	return MyMatrixBigInt{new(big.Int).Mul(a.Int, b.Int)}
}

// ================================================================================

type MatrixOperations[T any] interface {
	Add(T) T
	Mul(T) T
}

type Matrix2x2[T MatrixOperations[T]] [2][2]T

func (a Matrix2x2[T]) MulMatrix(b Matrix2x2[T]) Matrix2x2[T] {
	return Matrix2x2[T]{
		{
			a[0][0].Mul(b[0][0]).Add(a[0][1].Mul(b[1][0])),
			a[0][0].Mul(b[0][1]).Add(a[0][1].Mul(b[1][1])),
		},
		{
			a[1][0].Mul(b[0][0]).Add(a[1][1].Mul(b[1][0])),
			a[1][0].Mul(b[0][1]).Add(a[1][1].Mul(b[1][1])),
		},
	}
}

// ================================================================================

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	acc := Matrix2x2[MyMatrixInt]{
		{1, 1},
		{1, 0},
	}

	base := Matrix2x2[MyMatrixInt]{
		{1, 1},
		{1, 0},
	}

	for n > 0 {
		if n%2 != 0 {
			acc = acc.MulMatrix(base)
		}

		base = base.MulMatrix(base)
		n /= 2
	}

	return int(acc[0][1])
}
*/
