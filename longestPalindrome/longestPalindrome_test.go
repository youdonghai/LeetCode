package longestPalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	s := LongestPalindrome("c")
	if s != "c" {
		t.Error("longest palindrome of \"c\" is not c")
	}
	s = LongestPalindrome("cc")
	if s != "cc" {
		t.Error("longest palindrome of \"cc\" is not cc")
	}
	s = LongestPalindrome("aca")
	if s != "aca" {
		t.Error("longest palindrome of \"aca\" is not aca, it's ", s)
	}
	s = LongestPalindrome("hdxlxdgferiwo")
	if s != "dxlxd" {
		t.Error("longest palindrome of \"hdxlxdgferiwo\" is not dxlxd, it's ", s)
	}
}
