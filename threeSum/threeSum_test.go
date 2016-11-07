package threeSum

import "testing"

var er1 [][]int = [][]int{{-1, -1, 2}, {-1, 0, 1}}
var er2 [][]int = [][]int{}
var er3 [][]int = [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}}

func TestThreeSum(t *testing.T) {
	r := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	if len(r) != 2 {
		t.Error("test failed")
	}
	for i := 0; i < len(r); i++ {
		for j := 0; j < 3; j++ {
			if r[i][j] != er1[i][j] {
				t.Error("test failed", r, i, j)
			}
		}
	}
}

func TestThreeSum2(t *testing.T) {
	r := ThreeSum2([]int{-1, 0, 1, 2, -1, -4})
	if len(r) != 2 {
		t.Error("test failed")
	}
	for i := 0; i < len(r); i++ {
		for j := 0; j < 3; j++ {
			if r[i][j] != er1[i][j] {
				t.Error("test failed", r, i, j)
			}
		}
	}

	r = ThreeSum2([]int{0, 0})
	if len(r) != len(er2) {
		t.Error("test failed")
	}

	r = ThreeSum2([]int{3, 0, -2, -1, 1, 2})
	if len(r) != len(er3) {
		t.Error("test failed", r)
	}
	for i := 0; i < len(r); i++ {
		for j := 0; j < 3; j++ {
			if r[i][j] != er3[i][j] {
				t.Error("test failed", r, i, j)
			}
		}
	}
}
