package threeSum

import "testing"

func TestThreeSum(t *testing.T) {
	r := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	if len(r) != 2 {
		t.Error("test failed")
	}
	er := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	for i := 0; i < len(r); i++ {
		for j := 0; j < 3; j++ {
			if r[i][j] != er[i][j] {
				t.Error("test failed", r, i, j)
			}
		}
	}
}
