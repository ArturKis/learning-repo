package main

/*
# 206. Reverse Linked List

Given the head of a singly linked list, reverse the list, and return the reversed list.

Constraints:
  - The number of nodes in the list is the range [0, 5000].
  - -5000 <= Node.val <= 5000

Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
*/
func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/reverse-linked-list/description/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	/*
	 newHead > ........
	 ...n > .....
	*/
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

/*
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
// */
