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

type Node struct {
	val  int
	next *Node
	prev *Node
}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func Constructor() *MyLinkedList {
	return &MyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (this *MyLinkedList) FindNode(idx int) *Node {
	var curr *Node
	if idx < this.size/2 {
		curr = this.head

		for idx > 0 {
			curr = curr.next
			idx--
		}
	} else {
		curr = this.tail

		for idx < this.size-1 {
			curr = curr.prev
			idx++
		}
	}

	return curr
}

func (this *MyLinkedList) Get(idx int) int {
	if idx < 0 || idx >= this.size {
		return -1
	}

	return this.FindNode(idx).val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{val: val}

	if this.size == 0 {
		this.head = node
		this.tail = node
	} else {
		node.next = this.head
		this.head.prev = node
		this.head = node
	}

	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.size == 0 {
		this.AddAtHead(val)
	} else {
		node := &Node{val: val}

		node.prev = this.tail
		this.tail.next = node
		this.tail = node
		this.size++
	}
}

func (this *MyLinkedList) AddAtIndex(idx int, val int) {
	if idx < 0 || idx > this.size {
		return
	}

	switch idx {
	case 0:
		this.AddAtHead(val)
	case this.size:
		this.AddAtTail(val)
	default:
		node := &Node{val: val}
		curr := this.FindNode(idx)

		node.next = curr
		node.prev = curr.prev
		curr.prev.next = node
		curr.prev = node

		this.size++
	}
}

func (this *MyLinkedList) DeleteHead() {
	switch this.size {
	case 0:
		return
	case 1:
		this.head = nil
		this.tail = nil
		this.size--
	default:
		this.head = this.head.next
		this.head.prev.next = nil
		this.head.prev = nil
		this.size--
	}
}

func (this *MyLinkedList) DeleteTail() {
	switch this.size {
	case 0:
		return
	case 1:
		this.tail = nil
		this.head = nil
		this.size--
	default:
		this.tail = this.tail.prev
		this.tail.next.prev = nil
		this.tail.next = nil
		this.size--
	}
}

func (this *MyLinkedList) DeleteAtIndex(idx int) {
	if idx < 0 || idx >= this.size {
		return
	}

	switch idx {
	case 0:
		this.DeleteHead()
	case this.size - 1:
		this.DeleteTail()
	default:
		curr := this.FindNode(idx)

		curr.prev.next = curr.next
		curr.next.prev = curr.prev
		curr.prev, curr.next = nil, nil
		this.size--
	}
}

/* first implementation
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
*/
