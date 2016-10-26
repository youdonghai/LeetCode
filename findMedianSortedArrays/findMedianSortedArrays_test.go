package findMedianSortedArrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	r := FindMedianSortedArrays([]int{1, 3}, []int{2})
	if r != 2.0 {
		t.Error("find median of [1, 3] and [2] is not 2.0")
	}
	r = FindMedianSortedArrays([]int{1, 2}, []int{3, 4})
	if r != 2.5 {
		t.Error("find median of [1, 2] and [3, 4] is not 2.5")
	}
}
