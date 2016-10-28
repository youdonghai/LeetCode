package stringToInteger

import "testing"

func TestMyAtoi(t *testing.T) {
	r := MyAtoi("   --1231547395438")
	if r != 0 {
		t.Error("atoi test, two minus sign faild")
	}
	r = MyAtoi("+1231547")
	if r != 1231547 {
		t.Error("atoi test one plus sign faild")
	}
	r = MyAtoi("    5 4 3 8 9 0   ")
	if r != 5 {
		t.Error("atoi test faild", r)
	}
	r = MyAtoi("72189d13")
	if r != 72189 {
		t.Error("atoi test faild")
	}
	r = MyAtoi("9223372036854775809")
	if r != (1<<31)-1 {
		t.Error("atoi test faild", r)
	}
	var r32 int32
	r32 = 1
	r32 = r32 << 31
	min_int := int(r32)
	r = MyAtoi("-41435843950843905")
	if r != min_int {
		t.Error("atoi test faild")
	}
}
