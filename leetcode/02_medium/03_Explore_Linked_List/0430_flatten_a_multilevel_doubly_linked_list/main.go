package main

/*
# 430. Flatten a Multilevel Doubly Linked List

You are given a doubly linked list, which contains nodes that have a next pointer, a previous pointer, and an additional child pointer. This child pointer may or may not point to a separate doubly linked list, also containing these special nodes. These child lists may have one or more children of their own, and so on, to produce a multilevel data structure as shown in the example below.

Given the head of the first level of the list, flatten the list so that all the nodes appear in a single-level, doubly linked list. Let curr be a node with a child list. The nodes in the child list should appear after curr and before curr.next in the flattened list.

Return the head of the flattened list. The nodes in the list must have all of their child pointers set to null.

Constraints:

  - The number of Nodes will not exceed 1000.
  - 1 <= Node.val <= 10^5
*/
func main() {}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

// this made me angry | TODO: repeat later
// https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/description/
func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	var rec func(*Node) *Node
	rec = func(node *Node) *Node {
		curr := node
		for curr.Next != nil {
			if curr.Child != nil {
				next := curr.Next
				tail := rec(curr.Child)
				curr.Next = curr.Child
				curr.Child = nil
				curr.Next.Prev = curr
				tail.Next = next

				if next != nil {
					next.Prev = tail
				}

				curr = tail
			}

			curr = curr.Next
		}

		if curr.Child != nil {
			tail := rec(curr.Child)
			curr.Child.Prev = curr
			curr.Next = curr.Child
			curr.Child = nil
			curr = tail
		}

		return curr
	}

	rec(root)
	return root
}
