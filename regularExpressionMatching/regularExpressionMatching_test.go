package regularExpressionMatching

import "testing"

func TestIsMatch(t *testing.T) {
	r := IsMatch("", "")
	if r != true {
		t.Error("test empty string failed")
	}
	r = IsMatch("aa", "a")
	if r != false {
		t.Error("test aa and a failed")
	}
	r = IsMatch("aa", "aa")
	if r != true {
		t.Error("test aa and aa failed")
	}
	r = IsMatch("aaa", "aa")
	if r != false {
		t.Error("test aaa and aa failed")
	}
	r = IsMatch("aa", "a*")
	if r != true {
		t.Error("test aa and a* failed")
	}
	r = IsMatch("aa", ".*")
	if r != true {
		t.Error("test aa and .* failed")
	}
	r = IsMatch("aab", "c*a*b")
	if r != true {
		t.Error("test aab and c*a*b failed")
	}
	r = IsMatch("aaa", "a.a")
	if r != true {
		t.Error("test aaa and a.a failed")
	}
	r = IsMatch("aaa", "ab*ac*a")
	if r != true {
		t.Error("test aaa and ab*ac*a failed")
	}
	r = IsMatch("a", "ab*")
	if r != true {
		t.Error("test a and ab* failed")
	}
}
