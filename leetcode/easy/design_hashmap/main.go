package main

import "fmt"

/*
# Design HashMap

Design a HashMap without using any built-in hash table libraries.

Implement the MyHashMap class:

  - MyHashMap() initializes the object with an empty map.
  - void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
  - int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
  - void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.
*/
func main() {

	myHashMap := Constructor()

	myHashMap.Put(1, 1)           // The map is now [[1,1]]
	myHashMap.Put(2, 2)           // The map is now [[1,1], [2,2]]
	fmt.Println(myHashMap.Get(1)) // return 1, The map is now [[1,1], [2,2]]
	fmt.Println(myHashMap.Get(3)) // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
	myHashMap.Put(2, 1)           // The map is now [[1,1], [2,1]] (i.e., update the existing value)
	fmt.Println(myHashMap.Get(2)) // return 1, The map is now [[1,1], [2,1]]
	myHashMap.Remove(2)           // remove the mapping for 2, The map is now [[1,1]]
	fmt.Println(myHashMap.Get(2))
}

// https://leetcode.com/problems/design-hashmap/description/
type Entry struct {
	key   int
	value int
}

type MyHashMap struct {
	capacity int
	buckets  [][]Entry
}

func Constructor() MyHashMap {
	initCapacity := 100

	return MyHashMap{
		capacity: initCapacity,
		buckets:  make([][]Entry, initCapacity),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	bucketIdx := key % this.capacity

	for i, entry := range this.buckets[bucketIdx] {
		if entry.key == key {
			this.buckets[bucketIdx][i].value = value

			return
		}
	}

	this.buckets[bucketIdx] = append(this.buckets[bucketIdx], Entry{key, value})
}

func (this *MyHashMap) Get(key int) int {
	for _, v := range this.buckets[key%this.capacity] {
		if v.key == key {
			return v.value
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
	bucketIdx := key % this.capacity

	for i, v := range this.buckets[bucketIdx] {
		if v.key == key {
			bucket := this.buckets[bucketIdx]
			length := len(bucket)

			if length == 1 {
				this.buckets[bucketIdx] = make([]Entry, 0)
			} else {
				bucket[i] = bucket[length-1]
				this.buckets[bucketIdx] = bucket[:length-1]
			}

			break
		}
	}
}
