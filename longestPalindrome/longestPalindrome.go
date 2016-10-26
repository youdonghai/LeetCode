package longestPalindrome

func LongestPalindrome(s string) string {
	longest := 0
	longest_start := -1
	longest_end := -1
	bs := []byte(s)
	for i := 0; i < len(bs); i++ {
		check_len := 0
		start := i - check_len
		end := i + check_len
		for start >= 0 && end < len(bs) {
			if bs[start] != bs[end] {
				break
			}
			check_len++
			start = i - check_len
			end = i + check_len
		}
		if check_len*2-1 > longest {
			longest = check_len*2 - 1
			longest_start = start + 1
			longest_end = end
		}
		// another check
		check_len = 1
		start = i - check_len
		end = i - 1 + check_len
		for start >= 0 && end < len(bs) {
			if bs[start] != bs[end] {
				break
			}
			check_len++
			start = i - check_len
			end = i - 1 + check_len
		}
		if (check_len-1)*2 > longest {
			longest = (check_len - 1) * 2
			longest_start = start + 1
			longest_end = end
		}
	}
	return string(bs[longest_start:longest_end])
}
