package integerToRoman

func IntToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}
	roman_c := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman_n := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	r := ""
	for i := 0; i < len(roman_c); i++ {
		for t := num / roman_n[i]; t > 0; t-- {
			r += roman_c[i]
		}
		num = num % roman_n[i]
	}
	return r
}
