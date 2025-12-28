package main

import "fmt"

/*
# 387. First Unique Character in a String

Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.

Constraints:

  - 1 <= s.length <= 10^5
  - s consists of only lowercase English letters.
*/
func main() {
	testCases := []struct {
		s string
		e int
	}{
		{s: "leetcode", e: 0},
		{s: "loveleetcode", e: 2},
		{s: "aabb", e: -1},
		{s: "aabbc", e: 4},
		{s: "z", e: 0},
		{s: "dddccdbba", e: 8},
	}

	for i, tc := range testCases {
		result := firstUniqChar(tc.s)
		if result < 0 {
			fmt.Printf("Test %d (Input: %s => -1)\n", i+1, tc.s)
		} else {
			fmt.Printf("Test %d (Input: %s, Result: %d (%s))\n", i+1, tc.s, result, string(tc.s[result]))
		}
	}
}

// /*
// https://leetcode.com/problems/first-unique-character-in-a-string/description/
func firstUniqChar(s string) int {
	arr := [26]int{}

	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if arr[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

// */

/*
func firstUniqChar(s string) int {
	m := make(map[byte]int, 26)

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return i
		}
	}

	return -1
}
*/

/*
	func firstUniqChar(s string) int {
		m := make(map[rune]int)

		for _, v := range s {
			m[v] += 1
		}

		for i, v := range s {
			if m[v] == 1 {
				return i
			}
		}

		return -1
	}
*/
