package main

/*
# 328. Odd Even Linked List

Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

The first node is considered odd, and the second node is even, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in O(1) extra space complexity and O(n) time complexity.

Constraints:

  - The number of nodes in the linked list is in the range [0, 10^4].
  - -10^6 <= Node.val <= 10^6
*/
func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/odd-even-linked-list/submissions/1881106370/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr, evenHead := head, head.Next
	currEven := evenHead

	for currEven != nil && currEven.Next != nil {
		curr.Next = currEven.Next
		curr = curr.Next

		currEven.Next = curr.Next
		currEven = currEven.Next
	}

	curr.Next = evenHead

	return head
}
