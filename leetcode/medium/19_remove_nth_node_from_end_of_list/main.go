package main

/*
# 19. Remove Nth Node From End of List

Given the head of a linked list, remove the nth node from the end of the list and return its head.

Constraints:

  - The number of nodes in the list is sz.
  - 1 <= sz <= 30
  - 0 <= Node.val <= 100
  - 1 <= n <= sz

Follow up: Could you do this in one pass?
*/
func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	follow, lead := dummy, dummy

	for range n {
		if lead.Next == nil {
			return head.Next
		}

		lead = lead.Next
	}

	for lead.Next != nil {
		follow = follow.Next
		lead = lead.Next
	}

	orphan := follow.Next
	follow.Next = follow.Next.Next
	orphan.Next = nil

	return dummy.Next
}

/* first implementation
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	follow, lead := head, head

	for lead.Next != nil {
		if n == 0 {
			follow = follow.Next
		} else {
			n--
		}

		lead = lead.Next
	}

	if n == 0 {

		follow.Next.Next, follow.Next = nil, follow.Next.Next
	} else {
		head, head.Next = head.Next, nil
	}

	return head
}
*/
