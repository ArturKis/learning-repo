package main

/*
# 104. Maximum Depth of Binary Tree

Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Constraints:
  - The number of nodes in the tree is in the range [0, 10^4].
  - -100 <= Node.val <= 100
*/
func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ln := maxDepth(root.Left)
	rn := maxDepth(root.Right)

	return max(ln, rn) + 1
}
