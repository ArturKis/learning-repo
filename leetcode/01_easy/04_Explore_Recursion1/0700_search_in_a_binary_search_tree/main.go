package main

/*
# 700. Search in a Binary Search Tree

You are given the root of a binary search tree (BST) and an integer val.

Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.

Constraints:
  - The number of nodes in the tree is in the range [1, 5000].
  - 1 <= Node.val <= 10^7
  - root is a binary search tree.
  - 1 <= val <= 10^7
*/
func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/search-in-a-binary-search-tree/description/
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

/*
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	return root
}
*/
