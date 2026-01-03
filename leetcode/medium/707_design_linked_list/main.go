package main

import (
	"fmt"
)

/*
# 707. Design Linked List

Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and next is a pointer/reference to the next node.
If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

Implement the MyLinkedList class:

  - MyLinkedList() Initializes the MyLinkedList object.
  - int get(int index) Get the value of the indexth node in the linked list. If the index is invalid, return -1.
  - void addAtHead(int val) Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
  - void addAtTail(int val) Append a node of value val as the last element of the linked list.
  - void addAtIndex(int index, int val) Add a node of value val before the indexth node in the linked list. If index equals the length of the linked list, the node will be appended to the end of the linked list. If index is greater than the length, the node will not be inserted.
  - void deleteAtIndex(int index) Delete the indexth node in the linked list, if the index is valid.

Constraints:

  - 0 <= index, val <= 1000
  - Please do not use the built-in LinkedList library.
  - At most 2000 calls will be made to get, addAtHead, addAtTail, addAtIndex and deleteAtIndex.
*/
func main() {
	// "MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"
	// [[], [1], [3], [1, 2], [1], [1], [1]]
	ll := Constructor()

	ll.AddAtHead(1)
	ll.AddAtTail(3)
	ll.AddAtIndex(1, 2)
	fmt.Println(ll.Get(1))
	ll.DeleteAtIndex(1)
	fmt.Println(ll.Get(1))
}

// TODO: check tomorrow right implementation

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
}

// https://leetcode.com/problems/design-linked-list/description/
func Constructor() *MyLinkedList {
	return &MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if this.head == nil {
		return -1
	}

	curr := this.head

	for i := index; i > 0; i-- {
		if curr.next == nil {
			return -1
		}

		curr = curr.next
	}

	return curr.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	n := &Node{val: val, next: this.head}
	this.head = n
}

func (this *MyLinkedList) AddAtTail(val int) {
	n := &Node{val: val, next: nil}

	if this.head == nil {
		this.head = n
		return
	}

	curr := this.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = n
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if this.head == nil {
		return
	}

	curr := this.head
	for index > 1 {
		if curr.next == nil {
			return
		}

		curr = curr.next
		index--
	}

	n := &Node{val: val, next: nil}
	n.next = curr.next
	curr.next = n
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.head == nil {
		return
	}

	if index == 0 {
		this.head = this.head.next
		return
	}

	curr := this.head
	for index > 1 {
		if curr.next == nil {
			return
		}

		curr = curr.next
		index--
	}

	if curr.next == nil {
		return
	}

	curr.next, curr.next.next = curr.next.next, nil
}
