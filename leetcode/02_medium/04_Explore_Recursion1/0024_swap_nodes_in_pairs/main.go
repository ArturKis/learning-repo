package main

/*
# 24. Swap Nodes in Pairs

Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed.)

Constraints:
  - The number of nodes in the list is in the range [0, 100].
  - 0 <= Node.val <= 100
*/
func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}

/*
// https://leetcode.com/problems/swap-nodes-in-pairs/description/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	follow := dummy
	lead := head.Next

	for lead != nil {
		follow.Next.Next = lead.Next
		lead.Next = follow.Next
		follow.Next = lead

		if lead.Next.Next == nil {
			break
		}

		lead = lead.Next.Next.Next
		follow = follow.Next.Next
	}

	return dummy.Next
}
*/
