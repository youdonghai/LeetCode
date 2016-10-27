package reverseInteger

func Reverse(x int) int {
	negative := false
	if x < 0 {
		negative = true
		x = -x
	}
	r := 0
	for x > 0 {
		r *= 10
		r += x % 10
		x /= 10
	}
	if negative {
		r = -r
	}
	var r32 int32
	r32 = 1
	r32 = r32 << 31
	min_int := int(r32)
	if r > (1<<31)-1 || r < min_int {
		return 0
	}
	return r
}
