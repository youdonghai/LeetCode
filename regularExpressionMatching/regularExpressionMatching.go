package regularExpressionMatching

func IsMatch(s string, p string) bool {
	return isMatch(s, p)
}

func isMatch(s string, p string) bool {
	s_len := len(s)
	p_len := len(p)
	if p_len == 0 {
		if s_len == 0 {
			return true
		}
		return false
	}
	if p_len == 1 {
		if s_len != 1 {
			return false
		}
		if p[0] == '.' || p[0] == s[0] {
			return true
		}
		return false
	}
	if p[1] == '*' {
		for i := 0; i < s_len+1; i++ {
			if p[0] != '.' && i < s_len && s[i] != p[0] {
				return isMatch(s[i:], p[2:])
			}
			if isMatch(s[i:], p[2:]) {
				return true
			}
		}
	}
	if s_len < 1 {
		return false
	}
	if p[0] != '.' && s[0] != p[0] {
		return false
	}
	return isMatch(s[1:], p[1:])
}
