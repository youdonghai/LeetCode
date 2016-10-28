package stringToInteger

func MyAtoi(str string) int {
	sign := 0
	digital_count := 0
	result := 0
	for _, c := range str {
		if c == ' ' {
			if digital_count > 0 || sign > 0 {
				break
			}
			continue
		} else if c == '-' {
			if sign > 0 || digital_count > 0 {
				// multiple minus or plus sign
				return 0
			}
			sign = 1
		} else if c == '+' {
			if sign > 0 || digital_count > 0 {
				return 0
			}
			sign = 2
		} else if c >= '0' && c <= '9' {
			if sign == 0 && digital_count == 0 {
				sign = 2
			}
			digital_count++
			result = result*10 + int(byte(c)-'0')
			if result > (1 << 32) {
				break
			}
		} else {
			break
		}
	}
	if sign == 1 {
		result = -result
	}
	max_int := (1 << 31) - 1
	if result > max_int {
		result = max_int
	}
	var r32 int32
	r32 = 1
	r32 = r32 << 31
	min_int := int(r32)
	if result < min_int {
		result = min_int
	}
	return result
}
