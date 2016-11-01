package palindromeNumber

import "testing"

func TestIsPalindrome(t *testing.T) {
	r := IsPalindrome(0)
	if r != true {
		t.Error("0 test failed")
	}
	r = IsPalindrome(51015)
	if r != true {
		t.Error("51015 test failed")
	}
	r = IsPalindrome(543345)
	if r != true {
		t.Error("543345 test failed")
	}
	r = IsPalindrome(54)
	if r != false {
		t.Error("54 test failed")
	}
	r = IsPalindrome(10)
	if r != false {
		t.Error("10 test failed")
	}
	r = IsPalindrome(1000021)
	if r != false {
		t.Error("1000021 test failed")
	}
	r = IsPalindrome(-1)
	if r != false {
		t.Error("-1 test failed")
	}
}
