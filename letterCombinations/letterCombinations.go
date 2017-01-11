package letterCombinations

func LetterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return []string{}
	}
	d := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	l := []string{}
	for i := 0; i < len(digits); i++ {
		l = append(l, d[digits[i]])
	}
	combinations_len := 1
	index_mode_list := make([]int, len(l))
	for i := len(l) - 1; i >= 0; i-- {
		index_mode_list[i] = combinations_len
		combinations_len *= len(l[i])
	}
	r := []string{}
	tmp_str := make([]byte, len(l))
	for i := 0; i < combinations_len; i++ {
		for j := 0; j < len(l); j++ {
			tmp_str[j] = l[j][i/index_mode_list[j]%len(l[j])]
		}
		r = append(r, string(tmp_str))
	}
	return r
}
