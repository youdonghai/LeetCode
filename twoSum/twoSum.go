package twoSum

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		v, ok := m[complement]
		if ok && v != i {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return []int{-1, -1}
}
