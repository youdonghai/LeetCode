package containerWithMostWater

import "testing"

func TestMaxArea(t *testing.T) {
	r := MaxArea([]int{5, 6, 8, 4, 8, 0, 4, 6})
	if r != 36 {
		t.Error("test failed")
	}
}
