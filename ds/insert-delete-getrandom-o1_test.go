package ds

import (
	"fmt"
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	rs := NewRandomizedSet()
	elements := []int{1, 2, 3, 4, 5}
	for _, e := range elements {
		rs.Insert(e)
	}

	counts := make(map[int]int)
	numTrials := 100000

	for i := 0; i < numTrials; i++ {
		randElem := rs.GetRandom()
		counts[randElem]++
	}

	for _, e := range elements {
		if count, exists := counts[e]; exists {
			fmt.Printf("Element %d: %f%%\n", e, float64(count)/float64(numTrials)*100)
		} else {
			t.Errorf("Element %d was never selected by GetRandom", e)
		}
	}

	// Ensure the distribution is roughly uniform
	expectedFrequency := float64(numTrials) / float64(len(elements))
	tolerance := 0.05 * expectedFrequency // 5% tolerance

	for _, e := range elements {
		if count, exists := counts[e]; exists {
			if float64(count) < expectedFrequency-tolerance || float64(count) > expectedFrequency+tolerance {
				t.Errorf("Element %d has an unexpected frequency: %d", e, count)
			}
		}
	}
}
