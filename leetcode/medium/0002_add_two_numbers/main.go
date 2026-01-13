package main

/*
#2. Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Constraints:

	-The number of nodes in each linked list is in the range [1, 100].
	-0 <= Node.val <= 9
	-It is guaranteed that the list represents a number that does not have leading zeros.
*/
func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		carry = sum / 10
	}

	return dummy.Next
}

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	var solve func(*ListNode, *ListNode, int) *ListNode
	solve = func(l1, l2 *ListNode, carry int) *ListNode {
		if l1 == nil && l2 == nil && carry == 0 {
			return nil
		}

		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		curr := &ListNode{Val: sum % 10}
		curr.Next = solve(l1, l2, sum/10)
		return curr
	}

	return solve(l1, l2, carry)
}
*/

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	var carry int
	for l1 != nil || l2 != nil {
		curr.Next = &ListNode{}
		curr = curr.Next
		sum := 0

		switch true {
		case l1 != nil && l2 != nil:
			val1, val2 := l1.Val, l2.Val
			l1, l2 = l1.Next, l2.Next
			sum = val1 + val2 + carry
		case l1 != nil:
			val1 := l1.Val
			l1 = l1.Next
			sum = val1 + carry
		case l2 != nil:
			val2 := l2.Val
			l2 = l2.Next
			sum = val2 + carry
		}

		curr.Val, carry = sum%10, sum/10
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
*/
