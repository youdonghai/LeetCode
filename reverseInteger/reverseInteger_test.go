package reverseInteger

import "testing"

func TestReverse(t *testing.T) {
	r := Reverse(1234)
	if r != 4321 {
		t.Error("recerse 1234 failed")
	}
	r = Reverse(-410845)
	if r != -548014 {
		t.Error("recerse -410845 failed")
	}
	r = Reverse(-1)
	if r != -1 {
		t.Error("recerse -1 failed")
	}
	r = Reverse(0)
	if r != 0 {
		t.Error("recerse 0 failed")
	}
	r = Reverse(1534236469)
	if r != 0 {
		t.Error("recerse 0 failed")
	}
}
