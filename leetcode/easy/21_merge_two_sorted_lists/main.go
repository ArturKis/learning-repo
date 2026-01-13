package main

/*
# 21. Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Constraints:

  - The number of nodes in both lists is in the range [0, 50].
  - -100 <= Node.val <= 100
  - Both list1 and list2 are sorted in non-decreasing order.
*/
func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/merge-two-sorted-lists/description/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 == nil {
		curr.Next = list2
	} else {
		curr.Next = list1
	}

	return dummy.Next
}

/*
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var newHead *ListNode

	if list1.Val < list2.Val {
		newHead = list1
		newHead.Next = mergeTwoLists(list1.Next, list2)
	} else {
		newHead = list2
		newHead.Next = mergeTwoLists(list1, list2.Next)
	}

	return newHead
}
*/
