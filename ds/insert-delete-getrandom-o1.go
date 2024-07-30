package ds

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	nums    []int
	num2idx map[int]int
	rand    *rand.Rand
}

func NewRandomizedSet() *RandomizedSet {
	return &RandomizedSet{
		nums:    []int{},
		num2idx: make(map[int]int),
		rand:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Insert inserts a value to the set
// Returns true if the set did not already contain the specified element
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.num2idx[val]; ok {
		return false
	}
	this.num2idx[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

// Remove removes a value from the set
// Returns true if the set contained the specified element
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.num2idx[val]
	if !ok {
		return false
	}

	lastIdx := len(this.nums) - 1
	if idx < lastIdx {
		lastNum := this.nums[lastIdx]
		this.nums[idx] = lastNum
		this.num2idx[lastNum] = idx
	}
	this.nums = this.nums[:lastIdx]
	delete(this.num2idx, val)
	return true
}

// GetRandom get a random element from the set
func (this *RandomizedSet) GetRandom() int {
	return this.nums[this.rand.Intn(len(this.nums))]
}
