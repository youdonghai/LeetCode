package zigzagConversion

import "testing"

func TestConvert(t *testing.T) {
	r := Convert("", 1)
	if r != "" {
		t.Error("\"\" convert faild.")
	}
	r = Convert("d", 10)
	if r != "d" {
		t.Error("\"d\" convert faild.")
	}
	r = Convert("PAYPALISHIRING", 3)
	if r != "PAHNAPLSIIGYIR" {
		t.Error("\"PAYPALISHIRING\" convert faild.")
	}
	r = Convert("abcdefghijklmnopqrst", 6)
	if r != "akbjltcimsdhnregoqfp" {
		t.Error("\"abcdefghijklmnopqrst\" convert faild.")
	}
}
