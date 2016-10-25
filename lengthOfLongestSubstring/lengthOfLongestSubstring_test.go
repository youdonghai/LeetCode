package lengthOfLongestSubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	r := LengthOfLongestSubstring("abcabcbb")
	if r != 3 {
		t.Error("length of longest substing for \"abcabcbb\" is not 3")
	}
	r = LengthOfLongestSubstring("bbbbb")
	if r != 1 {
		t.Error("length of longest substing for \"bbbbb\" is not 1")
	}
	r = LengthOfLongestSubstring("pwwkew")
	if r != 3 {
		t.Error("length of longest substing for \"pwwkew\" is not 3")
	}
	r = LengthOfLongestSubstring("c")
	if r != 1 {
		t.Error("length of longest substing for \"c\" is not 1")
	}
}
