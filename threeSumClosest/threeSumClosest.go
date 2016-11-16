package threeSumClosest

func ThreeSumClosest(nums []int, target int) int {
	closest := (1 << 31) - 1
	r := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				v := nums[i] + nums[j] + nums[k]
				c := v - target
				if c < 0 {
					c = -c
				}
				if c < closest {
					closest = c
					r = v
				}
			}
		}
	}
	return r
}
