package median_of_sorted_arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedianOfSortedArrays(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		median float64
	}{
		{[]int{1}, []int{}, 1},
		{[]int{1, 2, 3, 4, 5}, []int{}, 3},
		{[]int{1, 3, 5}, []int{2, 4}, 3},
		{[]int{1, 2}, []int{3, 4, 5}, 3},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, test := range tests {
		result := findMedianSortedArrays(test.nums1, test.nums2)
		assert.Equal(t, test.median, result, "Median of %v and %v should be %v but got %v", test.nums1, test.nums2, test.median, result)
	}
}
