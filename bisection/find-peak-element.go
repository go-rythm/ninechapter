package bisection

func FindPeak(a []int) int {
	start, end := 1, len(a)-2
	for start+1 < end {
		mid := start + (end-start)/2
		if a[mid] < a[mid-1] {
			end = mid
		} else if a[mid] < a[mid+1] {
			start = mid
		} else {
			return mid
		}
	}
	if a[start] < a[end] {
		return end
	} else {
		return start
	}
}
