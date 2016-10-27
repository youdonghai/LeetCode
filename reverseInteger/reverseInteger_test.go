package reverseInteger

import "testing"

func TestReverse(t *testing.T) {
	r := Reverse(1234)
	if r != 4321 {
		t.Error("reverse 1234 failed")
	}
	r = Reverse(-410845)
	if r != -548014 {
		t.Error("reverse -410845 failed")
	}
	r = Reverse(-1)
	if r != -1 {
		t.Error("reverse -1 failed")
	}
	r = Reverse(0)
	if r != 0 {
		t.Error("reverse 0 failed")
	}
	r = Reverse(1534236469)
	if r != 0 {
		t.Error("reverse 1534236469 failed")
	}
}
