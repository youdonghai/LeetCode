package lengthOfLongestSubstring

func LengthOfLongestSubstring(s string) int {
	max_length := 0
	bs := []byte(s)
	alphabet_set := make(map[byte]int)
	for i := 0; i < len(bs); i++ {
		index, exist := alphabet_set[bs[i]]
		if exist {
			this_len := len(alphabet_set)
			if this_len > max_length {
				max_length = this_len
			}
			for k, v := range alphabet_set {
				if v <= index {
					delete(alphabet_set, k)
				}
			}
		}
		alphabet_set[bs[i]] = i
	}
	this_len := len(alphabet_set)
	if this_len > max_length {
		max_length = this_len
	}
	return max_length
}
