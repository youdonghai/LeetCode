package integerToRoman

import "testing"

func TestIntToRoman(t *testing.T) {
	r := IntToRoman(3999)
	if r != "MMMCMXCIX" {
		t.Error("test failed", r)
	}
	r = IntToRoman(1)
	if r != "I" {
		t.Error("test failed", r)
	}
	r = IntToRoman(657)
	if r != "DCLVII" {
		t.Error("test failed", r)
	}
}
