package main

import "fmt"

/*
# 771. Jewels and Stones

You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have. Each character in stones is a type of stone you have. You want to know how many of the stones you have are also jewels.

Letters are case sensitive, so "a" is considered a different type of stone from "A".

Constraints:

  - 1 <= jewels.length, stones.length <= 50
  - jewels and stones consist of only English letters.
  - All the characters of jewels are unique.
*/
func main() {
	// r1 := 'z'
	// r2 := 'A'
	// fmt.Println((r1 - r2))

	testCases := []struct {
		j string
		s string
		e int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for i, tc := range testCases {
		result := numJewelsInStones(tc.j, tc.s)
		fmt.Printf("Test #%d\nInput: {jewels: %s | stones: %s}\nExpected: %d (result: %d)\n\n", i+1, tc.j, tc.s, tc.e, result)
	}
}

// https://leetcode.com/problems/jewels-and-stones/description/
func numJewelsInStones(jewels string, stones string) int {
	jewelSet := 0 // 64bit
	count := 0

	for _, v := range jewels {
		jewelSet |= 1 << (v - 'A')
	}

	for _, v := range stones {
		if jewelSet&(1<<(v-'A')) != 0 {
			count++
		}
	}

	return count
}
