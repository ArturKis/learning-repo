package main

import (
	"fmt"
	"math/rand"
)

/*
# 380. Insert Delete GetRandom O(1)

Implement the RandomizedSet class:

  - RandomizedSet() Initializes the RandomizedSet object.
  - bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
  - bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
  - int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.

You must implement the functions of the class such that each function works in average O(1) time complexity.

Constraints:

  - -2^31 <= val <= 2^31 - 1
  - At most 2 * 10^5 calls will be made to insert, remove, and getRandom.
  - There will be at least one element in the data structure when getRandom is called.
*/
func main() {
	/*
		Input:
		["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
		[[], [1], [2], [2], [], [1], [2], []]
		Output:
		[null, true, false, true, 2, true, false, 2]
	*/

	test := Constructor()
	s := fmt.Sprintln(test.Insert(1), test.Remove(2), test.Insert(2), test.GetRandom(), test.Remove(1), test.Insert(2), test.GetRandom()) // expected output: [true false true randomInt true false randomInt]

	fmt.Println(s)
}

type RandomizedSet struct {
	bucket map[int]int
	slice  []int
}

// https://leetcode.com/problems/insert-delete-getrandom-o1/description/
func Constructor() RandomizedSet {
	return RandomizedSet{
		bucket: make(map[int]int),
		slice:  make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.bucket[val]; ok {
		return false
	}
	this.slice = append(this.slice, val)
	this.bucket[val] = len(this.slice) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.bucket[val]; ok {
		lastIdx := len(this.slice) - 1

		this.bucket[int(this.slice[lastIdx])] = idx

		this.slice[idx] = this.slice[lastIdx]
		this.slice = this.slice[:lastIdx]

		delete(this.bucket, val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	return int(this.slice[rand.Intn(len(this.slice))])
}
