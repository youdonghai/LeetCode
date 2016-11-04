package romanToInteger

import "testing"

func TestRomanToInt(t *testing.T) {
	r := RomanToInt("MMMCMXCIX")
	if r != 3999 {
		t.Error("test failed", r)
	}
	r = RomanToInt("I")
	if r != 1 {
		t.Error("test failed", r)
	}
	r = RomanToInt("DCLVII")
	if r != 657 {
		t.Error("test failed", r)
	}
}