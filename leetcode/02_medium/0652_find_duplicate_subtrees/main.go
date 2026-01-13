package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
# 652. Find Duplicate Subtrees

Given the root of a binary tree, return all duplicate subtrees.

For each kind of duplicate subtrees, you only need to return the root node of any one of them.

Two trees are duplicate if they have the same structure with the same node values.

Constraints:

  - The number of the nodes in the tree will be in the range [1, 5000]
  - -200 <= Node.val <= 200
*/
func main() {
	/*
		bt := []int{1, 2, 2, 3, -999_999_999, 4}
		slice := []*int{}

		for _, v := range bt {
			if v != -999_999_999 {
				slice = append(slice, intPtr(v))
			} else {
				slice = append(slice, nil)
			}
		}

		root := genNode(slice, 0)

		if root.Right.Left != nil {
			fmt.Println(root.Right.Left.Val)
		}

		fmt.Printf("Total nodes: %d\n", countNodes(root))
		// fmt.Printf("%v\n", treeToArr(root))
		fmt.Printf("%v\n", findDuplicateSubtrees(root))
	*/

	testCases := [][]int{
		{1, 2, 2, 3, -999_999_999, 4},                                   // 0
		{1, 2, 3, 4, -999_999_999, 2, 4, -999_999_999, -999_999_999, 4}, // [[2,4],[4]]  => 2
		{2, 1, 1}, // [[1]]  => 1
		{2, 2, 2, 3, -999_999_999, 3, -999_999_999}, // [[2,3],[3]]  => 2
	}

	for i, tc := range testCases {
		slice := []*int{}

		for _, v := range tc {
			if v != -999_999_999 {
				slice = append(slice, intPtr(v))
			} else {
				slice = append(slice, nil)
			}
		}

		root := genNode(slice, 0)

		cn := countNodes(root)
		result := len(findDuplicateSubtrees(root))

		fmt.Printf("Test %d\nTotal nodes: %d\nDuplicates: %v\n\n", i+1, cn, result)
	}
}

// https://leetcode.com/problems/find-duplicate-subtrees/description/
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	result := []*TreeNode{}
	m := make(map[string]int)

	var traverse func(*TreeNode) string
	traverse = func(n *TreeNode) string {
		if n == nil {
			return "#"
		}

		left := traverse(n.Left)
		right := traverse(n.Right)

		var sb strings.Builder
		sb.WriteString(strconv.Itoa(n.Val))
		sb.WriteString("|")
		sb.WriteString(left)
		sb.WriteString("|")
		sb.WriteString(right)
		sb.WriteString("|")

		s := sb.String()

		m[s]++
		if m[s] == 2 {
			result = append(result, n)
		}

		return s
	}

	traverse(root)

	return result
}

// first look on binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func intPtr(n int) *int {
	return &n
}

func genNode(s []*int, idx int) *TreeNode {
	if idx >= len(s) || s[idx] == nil {
		return nil
	}

	return &TreeNode{Val: *s[idx], Left: genNode(s, idx*2+1), Right: genNode(s, idx*2+2)}
}

func countNodes(n *TreeNode) int {
	if n == nil {
		return 0
	}

	return 1 + countNodes(n.Left) + countNodes(n.Right)
}

func treeToArr(n *TreeNode) []*TreeNode {
	if n == nil {
		return []*TreeNode{}
	}

	arr := []*TreeNode{n}
	arr = append(arr, treeToArr(n.Left)...)
	arr = append(arr, treeToArr(n.Right)...)

	return arr
}

/*
         1
    2         2
   / \       / \
  3   nil   4  nil
*/
