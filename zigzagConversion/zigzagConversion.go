package zigzagConversion

func Convert(s string, numRows int) string {
	zigzag := make(map[int]string)
	now_row := 0
	at_mid := false
	add_value := 1
	for i := 0; i < len(s); i++ {
		if now_row == -1 {
			now_row++
			add_value = 1
			at_mid = false
		} else if now_row == numRows {
			now_row--
			add_value = -1
			at_mid = true
		}
		if at_mid && (now_row == 0 || now_row == numRows-1) {
			i--
			now_row += add_value
			continue
		}
		zigzag[now_row] += s[i : i+1]
		now_row += add_value
	}
	r := ""
	for i := 0; i < numRows; i++ {
		r += zigzag[i]
	}
	return r
}
