package bisection

import "sort"

/**
 * @param l: Given n pieces of wood with length L[i]
 * @param k: An integer
 * @return: The maximum length of the small pieces
 */
func WoodCut(l []int, k int) int {
	if len(l) == 0 {
		return 0
	}
	sort.Ints(l)
	start, end := 1, l[len(l)-1]
	for start+1 < end {
		mid := start + (end-start)/2
		if getPieces(l, mid) >= k {
			start = mid
		} else {
			end = mid
		}
	}

	if getPieces(l, end) >= k {
		return end
	}
	if getPieces(l, start) >= k {
		return start
	}
	return 0
}

func getPieces(l []int, length int) int {
	var pieces int
	for _, v := range l {
		pieces += v / length
	}
	return pieces
}
