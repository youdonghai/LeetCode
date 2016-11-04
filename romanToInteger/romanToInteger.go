package romanToInteger

func RomanToInt(s string) int {
	r := 0
	roman_c := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman_n := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for len(s) > 0 {
		for i := 0; i < 13; i++ {
			c_len := len(roman_c[i])
			if len(s) < c_len {
				continue
			}
			if s[0:c_len] == roman_c[i] {
				r += roman_n[i]
				s = s[c_len:]
				break
			}
		}
	}
	return r
}
