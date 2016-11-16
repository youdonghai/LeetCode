package threeSumClosest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	r := ThreeSumClosest([]int{0, 0, 0, 1, 6, 7, 8, 4, -76, 3, -54, 7, 3, 5, 7, 8, 3, 6, 7, 3, 7, 5, 4}, 99)
	if r != 23 {
		t.Error("test error")
	}
}
