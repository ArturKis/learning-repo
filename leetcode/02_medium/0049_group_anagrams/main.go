package main

import "fmt"

/*
# 49. Group Anagrams

Given an array of strings strs, group the anagrams together. You can return the answer in any order.

Constraints:

  - 1 <= strs.length <= 104
  - 0 <= strs[i].length <= 100
  - strs[i] consists of lowercase English letter
*/
func main() {
	testCases := []struct {
		i []string
		e [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
	}

	for i, tc := range testCases {
		result := groupAnagrams(tc.i)
		fmt.Printf("Test #%d\nInput: %v\nResult: %v\nExpected: %v\n", i+1, tc.i, result, tc.e)
	}

}

// https://leetcode.com/problems/group-anagrams/description/
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, v := range strs {
		mask := [26]int{}
		for _, r := range v {
			mask[r-'a']++
		}
		m[mask] = append(m[mask], v)
	}

	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

/* brute force
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	result := [][]string{}

	for _, v := range strs {
		r := []rune(v)
		slices.Sort(r)
		s := string(r)

		m[s] = append(m[s], v)
	}

	for _, v := range m {
		result = append(result, v)
	}

	return result
}
*/
