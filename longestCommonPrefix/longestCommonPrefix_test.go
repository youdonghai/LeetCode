package longestCommonPrefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	r := LongestCommonPrefix([]string{"abc", "ab", "a"})
	if r != "a" {
		t.Error("test failed", r)
	}
	r = LongestCommonPrefix([]string{"abc", "ab", ""})
	if r != "" {
		t.Error("test failed", r)
	}
	r = LongestCommonPrefix([]string{})
	if r != "" {
		t.Error("test failed", r)
	}
}
