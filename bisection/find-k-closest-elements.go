package bisection

/**
 * @param a: an integer array
 * @param target: An integer
 * @param k: An integer
 * @return: an integer array
 */
func KClosestNumbers(a []int, target int, k int) []int {
	left := findLowerClosest(a, target)
	right := left + 1

	res := make([]int, k)
	for i := 0; i < k; i++ {
		if isLeftCloser(a, target, left, right) {
			res[i] = a[left]
			left--
		} else {
			res[i] = a[right]
			right++
		}
	}

	return res
}

func isLeftCloser(a []int, target int, left, right int) bool {
	if left < 0 {
		return false
	}
	if right >= len(a) {
		return true
	}
	if target-a[left] != a[right]-target {
		return target-a[left] < a[right]-target
	}
	return true
}

func findLowerClosest(a []int, target int) int {
	start, end := 0, len(a)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if a[mid] < target {
			start = mid
		} else {
			end = mid
		}
	}

	if a[end] < target {
		return end
	}
	if a[start] < target {
		return start
	}
	return -1
}
