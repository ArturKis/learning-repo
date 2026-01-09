package main

/*
# 206. Reverse Linked List

Given the head of a singly linked list, reverse the list, and return the reversed list.

Constraints:

	The number of nodes in the list is the range [0, 5000].
	-5000 <= Node.val <= 5000

Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
*/
func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// /*
// https://leetcode.com/problems/reverse-linked-list/description/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	h := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return h
}

// */

/*
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}

	for head.Next != nil {
		n := head.Next
		head.Next = n.Next
		n.Next = dummy.Next
		dummy.Next = n
	}

	return dummy.Next
}
*/
