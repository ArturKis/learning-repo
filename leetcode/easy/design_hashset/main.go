package main

import (
	"fmt"
	"slices"
)

/*
# Design HashSet

Design a HashSet without using any built-in hash table libraries.

Implement MyHashSet class:

  - void add(key) Inserts the value key into the HashSet.

  - bool contains(key) Returns whether the value key exists in the HashSet or not.

  - void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.
*/
func main() {
	myHashSet := Constructor()
	myHashSet.Add(1)
	myHashSet.Add(2)
	fmt.Println(myHashSet.Contains(1))
	fmt.Println(myHashSet.Contains(3))
	myHashSet.Add(2)
	fmt.Println(myHashSet.Contains(2))
	myHashSet.Remove(2)
	fmt.Println(myHashSet.Contains(2))
}

// https://leetcode.com/problems/design-hashset/description/
type MyHashSet struct {
	buckets  [][]int
	capacity int
}

func Constructor() MyHashSet {
	initCap := 100

	return MyHashSet{
		buckets:  make([][]int, initCap),
		capacity: initCap,
	}
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		idx := key % this.capacity
		this.buckets[idx] = append(this.buckets[idx], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	idx := key % this.capacity
	bucket := this.buckets[idx]
	length := len(bucket)

	for i, v := range bucket {
		if v == key {
			bucket[i] = bucket[length-1]
			this.buckets[idx] = bucket[:length-1]
			break
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	// for _, v := range this.buckets[key%this.capacity] {
	// 	if v == key {
	// 		return true
	// 	}
	// }
	// return false

	return slices.Contains(this.buckets[key%this.capacity], key)
}
