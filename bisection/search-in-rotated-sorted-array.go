package bisection

/**
 * @param a: an integer rotated sorted array
 * @param target: an integer to be searched
 * @return: an integer
 */
func Search(a []int, target int) int {
	if len(a) == 0 {
		return -1
	}
	start, end := 0, len(a)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if a[mid] >= a[0] {
			if target >= a[0] && target <= a[mid] {
				end = mid
			} else {
				start = mid
			}
		} else {
			if target >= a[mid] && target <= a[end] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if a[start] == target {
		return start
	}
	if a[end] == target {
		return end
	}
	return -1
}
