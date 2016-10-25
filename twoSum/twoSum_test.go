package twoSum

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 5, 8, 7}
	target := 15
	r := TwoSum(nums, target)
	if len(r) != 2 || r[0] != 2 || r[1] != 3 {
		t.Error("test twoSum failed")
	} else {
		t.Log("test twoSum pass")
	}
}
