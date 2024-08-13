package heap

// Heapify transforms the array into a min-heap.
func Heapify(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		siftDown(nums, i)
	}
}

// siftDown adjusts the elements to maintain the min-heap property.
func siftDown(nums []int, index int) {
	n := len(nums)
	for index < n {
		left := index*2 + 1
		right := index*2 + 2
		minIndex := index

		if left < n && nums[left] < nums[minIndex] {
			minIndex = left
		}
		if right < n && nums[right] < nums[minIndex] {
			minIndex = right
		}

		if minIndex == index {
			break
		}

		nums[minIndex], nums[index] = nums[index], nums[minIndex]
		index = minIndex
	}
}
