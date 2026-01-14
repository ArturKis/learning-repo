package main

/*
# 138. Copy List with Random Pointer

A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

Return the head of the copied linked list.

The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

  - val: an integer representing Node.val
  - random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.

Your code will only be given the head of the original linked list.

Constraints:

  - 0 <= n <= 1000
  - -10^4 <= Node.val <= 10^4
  - Node.random is null or is pointing to some node in the linked list.
*/
func main() {}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
// recursion approach again made me angry | TODO: repeat later
func copyRandomList(head *Node) *Node {
	mapNodes := make(map[*Node]*Node)

	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}

		if _, ok := mapNodes[node]; !ok {
			mapNodes[node] = &Node{Val: node.Val}
			dfs(node.Next)
			dfs(node.Random)
		}

		mapNodes[node].Next = mapNodes[node.Next]
		mapNodes[node].Random = mapNodes[node.Random]
	}

	dfs(head)

	newHead := mapNodes[head]
	mapNodes = nil

	return newHead
}
*/

// /* implementation with map | memory: O(n) | time: O(n)  =>  looks better

// https://leetcode.com/problems/copy-list-with-random-pointer/submissions/1884424284/
func copyRandomList(head *Node) *Node {
	mapNodes := map[*Node]*Node{}

	curr := head

	for curr != nil {
		mapNodes[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	curr = head

	for curr != nil {
		mapNodes[curr].Next = mapNodes[curr.Next]
		mapNodes[curr].Random = mapNodes[curr.Random]

		curr = curr.Next
	}

	newHead := mapNodes[head]
	mapNodes = nil

	return newHead
}

// */

/*
// brute force O(n^2)
func copyRandomList(head *Node) *Node {
	indexes := []int{}

	curr1 := head

	for curr1 != nil {
		if curr1.Random != nil {
			curr2 := head

			for i := 0; curr2 != nil; i++ {
				if curr2 == curr1.Random {
					indexes = append(indexes, i)
				}
				curr2 = curr2.Next
			}
		} else {
			indexes = append(indexes, -1)
		}

		curr1 = curr1.Next
	}

	dummy := &Node{}
	curr := dummy
	curr1 = head

	for curr1 != nil {
		curr.Next = &Node{Val: curr1.Val}
		curr = curr.Next
		curr1 = curr1.Next
	}

	curr = dummy.Next

	for i := 0; curr != nil; i++ {
		if indexes[i] != -1 {
			if indexes[i]-i >= 0 {
				pos := indexes[i] - i
				curr2 := curr

				for pos > 0 {
					curr2 = curr2.Next
					pos--
				}

				curr.Random = curr2
			} else {
				pos := indexes[i]
				curr2 := dummy.Next

				for pos > 0 {
					curr2 = curr2.Next
					pos--
				}

				curr.Random = curr2
			}
		}

		curr = curr.Next
	}

	return dummy.Next
}
*/
