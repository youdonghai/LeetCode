package letterCombinations

import "testing"

func testEq(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestLetterCombinations(t *testing.T) {
	r := LetterCombinations("23")
	if !testEq(r, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}) {
		t.Error("\"23\" combinations error")
	}
	r = LetterCombinations("")
	if !testEq(r, []string{}) {
		t.Error("\"\" combinations error")
	}
}
