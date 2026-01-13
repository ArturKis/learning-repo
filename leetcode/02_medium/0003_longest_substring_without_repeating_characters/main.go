package main

import "fmt"

/*
# 3. Longest Substring Without Repeating Characters

Given a string s, find the length of the longest substring without duplicate characters.

Constraints:

  - 0 <= s.length <= 5 * 10^4
  - s consists of English letters, digits, symbols and spaces.
*/
func main() {
	testCases := []struct {
		i string
		e int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"dvdf", 3},
	}

	for i, tc := range testCases {
		result := lengthOfLongestSubstring(tc.i)
		fmt.Printf("Test #%d\nInput: %#v\nExpected: %d => result: %d\n\n", i+1, tc.i, tc.e, result)
	}

	// fmt.Println(' ', 'z', '/', ',')
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	arr := [128]int{}

	l := 0

	for r := 0; r < len(s); r++ {
		char := s[r]

		if arr[char] > l {
			l = arr[char]
		}

		maxLen = max(maxLen, r-l+1)

		arr[char] = r + 1
	}

	return maxLen
}

/* brute force
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	m := make(map[byte]struct{})

	l := 0
	r := 0

	for r < len(s) {
		cur := s[r]

		if _, ok := m[cur]; ok {
			for l < r {
				if _, stillExist := m[cur]; !stillExist {
					break
				}
				delete(m, s[l])
				l++
			}
		}

		m[cur] = struct{}{}
		maxLen = max(maxLen, r-l+1)
		r++
	}

	return maxLen
}
*/
