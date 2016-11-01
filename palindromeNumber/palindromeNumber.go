package palindromeNumber

import "math"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	max_decimal := 0
	for n := x; n >= 10; n /= 10 {
		max_decimal++
	}
	for max_decimal > 0 {
		if x/int(math.Pow(10, float64(max_decimal))) != x%10 {
			return false
		}
		x = (x - (x%10)*int(math.Pow(10, float64(max_decimal)))) / 10
		max_decimal -= 2
	}
	return true
}
