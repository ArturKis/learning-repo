package main

/*
# 234. Palindrome Linked List

	Given the head of a singly linked list, return true if it is a palindrome or false otherwise.

	Constraints:

	    - The number of nodes in the list is in the range [1, 10^5].
	    - 0 <= Node.val <= 9


	Follow up: Could you do it in O(n) time and O(1) space?
*/
func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/palindrome-linked-list/submissions/1881489799/
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	if head.Next.Next == nil {
		return head.Val == head.Next.Val
	}

	follow, lead := head, head

	for lead != nil && lead.Next != nil {
		lead = lead.Next.Next
		follow = follow.Next
	}

	curr := follow.Next
	prev := follow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	left := head
	result := true

	for left != follow {
		if prev.Val != left.Val && result {
			result = false
		}
		left = left.Next

		next := prev.Next
		prev.Next = curr
		curr = prev
		prev = next
	}

	return result
}

/*
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}

	var slice []int

	follow, lead := &ListNode{Next: head}, &ListNode{Next: head}

	for lead != nil && lead.Next != nil {
		follow = follow.Next
		slice = append(slice, follow.Val)
		lead = lead.Next.Next
	}

	for i := len(slice) - 1; i >= 0; i-- {
		if lead != nil {
			follow = follow.Next
			if follow.Val != slice[i] {
				return false
			}
		} else {
			if follow.Val != slice[i] {
				return false
			}
			follow = follow.Next
		}
	}

	return true
}
*/
