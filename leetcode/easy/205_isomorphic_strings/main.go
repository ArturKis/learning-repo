package main

import "fmt"

/*
# Isomorphic Strings

Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

Constraints:

  - 1 <= s.length <= 5 * 10^4
  - t.length == s.length
  - s and t consist of any valid ascii character.
*/
func main() {
	testCases := []struct {
		s string
		t string
		e bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"badc", "baba", false},
	}

	for _, tc := range testCases {
		fmt.Printf("isIsomorphic(%s, %s) = %v, want %v\n", tc.s, tc.t, isIsomorphic(tc.s, tc.t), tc.e)

	}
}

// https://leetcode.com/problems/isomorphic-strings/description/
func isIsomorphic(s string, t string) bool {
	sMap, tMap := make(map[byte]byte), make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sr, tr := s[i], t[i]

		if mv, ok := sMap[sr]; ok {
			if mv != tr {
				return false
			}
		} else {
			sMap[sr] = tr
		}

		if mv, ok := tMap[tr]; ok {
			if mv != sr {
				return false
			}
		} else {
			tMap[tr] = sr
		}
	}

	return true
}

/*
func isIsomorphic(s string, t string) bool {
	m := make(map[string]string)

	for i, v := range s {
		char1 := string(v)
		char2 := string(t[i])

		if mv, ok := m[char1]; ok {

			if mv != char2 {
				return false
			}

		} else {

			for _, mv := range m {
				if mv == char2 {
					return false
				}
			}

		}

		m[char1] = char2
	}

	return true
}
*/
