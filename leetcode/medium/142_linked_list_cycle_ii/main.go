package main

/*
# 142. Linked List Cycle II

Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.

Constraints:

  - The number of the nodes in the list is in the range [0, 10^4].
  - -10^5 <= Node.val <= 10^5
  - pos is -1 or a valid index in the linked-list.

Follow up: Can you solve it using O(1) (i.e. constant) memory?
*/
func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-cycle-ii/description/
func detectCycle(head *ListNode) *ListNode {
	slow, fast, result := head, head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			for slow != result {
				slow = slow.Next
				result = result.Next
			}

			return result
		}
	}

	return nil
}

//  0   1   2   3   4   5    6   7    8   9   10 -> 7 (11) loop 4
// 0|0 1|2 2|4 3|6 4|8 5|10 6|8 7|10 8|8
