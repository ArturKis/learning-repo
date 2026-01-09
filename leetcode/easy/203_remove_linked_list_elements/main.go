package main

/*
# 203. Remove Linked List Elements

Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

Constraints:

  - The number of nodes in the list is in the range [0, 10^4].
  - 1 <= Node.val <= 50
  - 0 <= val <= 50
*/
func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/remove-linked-list-elements/description/
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	var currNode = dummy

	for currNode.Next != nil {
		if n := currNode.Next; n.Val == val {
			currNode.Next = n.Next
			n.Next = nil
		} else {
			currNode = n
		}
	}

	return dummy.Next
}
