package main

import "fmt"

/*
# 344. Reverse String

Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

Constraints
  - 1 <= s.length <= 10^5
  - s[i] is a printable ascii character.
*/
func main() {
	testCases := []struct {
		input []byte
	}{
		{input: []byte{'h', 'e', 'l', 'l', 'o'}},
		{input: []byte{'H', 'a', 'n', 'n', 'a', 'h'}},
	}

	for _, tc := range testCases {
		reverseString(tc.input)
		fmt.Println(string(tc.input))
	}

	reverseString(testCases[0].input)
	fmt.Println(string(testCases[0].input))
}

// https://leetcode.com/problems/reverse-string/description/
func reverseString(s []byte) {
	var helper func(int, int)
	helper = func(left int, right int) {
		if right <= left {
			return
		}

		s[left], s[right] = s[right], s[left]
		left++
		right--

		helper(left, right)
	}

	helper(0, len(s)-1)
}

/*
func reverseString(s []byte) {
	sliceLen := len(s)
	if sliceLen < 1 {
		return
	}

	for l, r := 0, sliceLen-1; l <= r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}
*/
