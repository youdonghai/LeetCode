package containerWithMostWater

func MaxArea(height []int) int {
	max := 0
	for i, h1 := range height {
		for j, h2 := range height[i+1:] {
			shorter := h1
			if h1 > h2 {
				shorter = h2
			}
			c := shorter * (j + 1)
			if c > max {
				max = c
			}
		}
	}
	return max
}
