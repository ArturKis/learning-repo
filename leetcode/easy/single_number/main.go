package main

import "fmt"

/*
# Single Number

Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

@Constraints:

  - 1 <= nums.length <= 3 * 10^4
  - -3 * 10^4 <= nums[i] <= 3 * 10^4
  - Each element in the array appears twice except for one element which appears only once.
*/
func main() {
	// /*
	nums1 := []int{2, 2, 1}
	nums2 := []int{4, 1, 2, 1, 2}
	nums3 := []int{1}

	fmt.Printf("Slice Nums1: %d\n", singleNumber(nums1))
	fmt.Printf("Slice Nums2: %d\n", singleNumber(nums2))
	fmt.Printf("Slice Nums3: %d\n", singleNumber(nums3))

	// test TSN
	mtx := [][]int{
		{1, 2, 1, 3, 2, 5},
		{-1, 0, -1, 4},
		{2, 3},
		{1, 2, 1, 3, 2, 5},
		{-1, 0, -1, 4},
		{10, 10, 2, 3},
		{1000000, 2, 1000000, 500},
	}

	PrintTSN(mtx)
	// */

	// madNums1 := []int8{5, 3, 5, 5}
	// madNums2 := []int8{2, 5, 2, 3, 5, 2, 5}
	madNums3 := []int8{2, 5, 2, 10, -7, 5, 2, 5, -7, -7}
	// fmt.Println(annoyingSingleNumber(madNums1))
	// fmt.Println(annoyingSingleNumber(madNums2))
	fmt.Println(annoyingSingleNumber(madNums3))

}

// /* https://leetcode.com/problems/single-number/
func singleNumber(nums []int) int {
	number := 0

	for _, v := range nums {
		number ^= v
	}

	return number
}

// My solution if two single elements. For practice and learn more about bitwise
func twoSingleNumbers(nums []int) (int, int) {
	num1 := 0
	num2 := 0
	mask := 1

	stuckTogether := 0

	for _, v := range nums {
		stuckTogether ^= v
	}

	for (mask & stuckTogether) == 0 {
		mask <<= 1
	}

	for _, v := range nums {
		if v&mask != 0 {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}

	return num1, num2
}

// */

// /*
func PrintTSN(matrix [][]int) {
	for i, v := range matrix {
		num1, num2 := twoSingleNumbers(v)
		fmt.Printf("Slice %d, num1: %d | num2: %d\n", i+1, num1, num2)
	}
}

// */

// this make me mad, lost over 10 hours and mental health. !DID'T FIND SOLUTION!
func annoyingSingleNumber(nums []int8) int8 {
	var a int8 = 0
	var b int8 = 0
	// var с int8 = 0

	for i, v := range nums {
		fmt.Println("Iter:", i+1)
		fmt.Printf("v => %08b\n", uint8(v))
		fmt.Printf("a => %08b\n", uint8(a))
		fmt.Printf("b => %08b\n", uint8(b))
		fmt.Println("***")

		fmt.Println("a = (a ^ v) & ^b")
		fmt.Printf("(%08b ^ %08b) &  ^%08b\n", uint8(a), uint8(v), uint8(b))
		fmt.Printf("%08b &  %08b\n", uint8(a^v), uint8(^b))

		a = (a ^ v) & ^b
		fmt.Printf("a => %08b\n", uint8(a))

		fmt.Println("***")

		fmt.Println("b = (b ^ v) & ^a")
		fmt.Printf("(%08b ^ %08b) &  ^%08b\n", uint8(b), uint8(v), uint8(a))
		fmt.Printf("%08b &  %08b\n", uint8(b^v), uint8(^a))

		b = (b ^ v) & ^a
		fmt.Printf("b => %08b\n", uint8(b))

		fmt.Println("=====================")
	}

	// fmt.Println(a, b)
	// fmt.Printf("a: %08b | b: %08b \n", a, b)
	// fmt.Printf("a: %08b | b: %08b | three: %08b\n", a, b, с)
	return a
}
