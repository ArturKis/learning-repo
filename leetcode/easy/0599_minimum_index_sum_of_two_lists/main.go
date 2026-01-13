package main

import "fmt"

/*
# 599. Minimum Index Sum of Two Lists

Given two arrays of strings list1 and list2, find the common strings with the least index sum.

A common string is a string that appeared in both list1 and list2.

A common string with the least index sum is a common string such that if it appeared at list1[i] and list2[j] then i + j should be the minimum value among all the other common strings.

Return all the common strings with the least index sum. Return the answer in any order.

Constraints:

  - 1 <= list1.length, list2.length <= 1000
  - 1 <= list1[i].length, list2[i].length <= 30
  - list1[i] and list2[i] consist of spaces ' ' and English letters.
  - All the strings of list1 are unique.
  - All the strings of list2 are unique.
  - There is at least a common string between list1 and list2.
*/
func main() {
	testCases := []struct {
		l1 []string
		l2 []string
		e  []string
	}{
		{
			l1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			l2: []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			e:  []string{"Shogun"},
		},
		{
			l1: []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			l2: []string{"KFC", "Shogun", "Burger King"},
			e:  []string{"Shogun"},
		},
		{
			l1: []string{"happy", "sad", "good"},
			l2: []string{"sad", "happy", "good"},
			e:  []string{"sad", "happy"},
		},
		{
			l1: []string{"A", "B", "C"},
			l2: []string{"C", "B", "A"},
			e:  []string{"A", "B", "C"},
		},
		{
			l1: []string{"void", "test", "go"},
			l2: []string{"apple", "pear", "void"},
			e:  []string{"void"},
		},
	}

	for i, tc := range testCases {
		res := findRestaurant(tc.l1, tc.l2)
		fmt.Printf("Test %d:\n  Input 1: %v\n  Input 2: %v\n  Expected: %v\n  Result  : %v\n\n", i+1, tc.l1, tc.l2, tc.e, res)
	}
}

// https://leetcode.com/problems/minimum-index-sum-of-two-lists/description/
func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) > len(list2) {
		return findRestaurant(list2, list1)
	}

	m := make(map[string]int, len(list1))
	minSum := len(list1) + len(list2)
	result := []string{}

	for i, v := range list1 {
		m[v] = i
	}

	for i, v := range list2 {
		if mv, ok := m[v]; ok {
			idxSum := mv + i

			if idxSum < minSum {
				minSum = idxSum
				result = result[:0]
				result = append(result, v)
			} else if idxSum == minSum {
				result = append(result, v)
			}
		}
	}

	return result
}
