package longestCommonPrefix

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	p := strs[0]
	for _, s := range strs[1:] {
		if len(s) < len(p) {
			p = p[:len(s)]
		}
		for i := 0; i < len(p); i++ {
			if p[i] != s[i] {
				p = p[:i]
				break
			}
		}
	}
	return p
}
