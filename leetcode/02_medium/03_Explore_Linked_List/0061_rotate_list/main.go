package main

/*
# 61. Rotate List

Given the head of a linked list, rotate the list to the right by k places.

Constraints:

  - The number of nodes in the list is in the range [0, 500].
  - -100 <= Node.val <= 100
  - 0 <= k <= 2 * 10^9
*/
func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/rotate-list/description/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	dummy := &ListNode{Next: head}
	follow, lead := dummy, dummy

	count := 0

	for lead.Next != nil {
		lead = lead.Next
		count++
	}

	k %= count
	if k == 0 {
		return head
	}

	k = count - k
	for k > 0 {
		follow = follow.Next
		k--
	}

	newHead := follow.Next
	follow.Next = nil
	lead.Next = head

	return newHead
}
