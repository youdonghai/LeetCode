package threeSum

import "sort"
import "strconv"

func ThreeSum(nums []int) [][]int {
	// not red-black tree in golang, so ...
	r := [][]int{}
	two_sum_map := make(map[int][]int)
	n := len(nums)
	for i, v1 := range nums {
		unique_map := make(map[string]bool)
		for j, v2 := range nums[i+1:] {
			ustr := ""
			if v1 > v2 {
				ustr = strconv.Itoa(v2) + strconv.Itoa(v1)
			} else {
				ustr = strconv.Itoa(v1) + strconv.Itoa(v2)
			}
			_, exist := unique_map[ustr]
			if exist {
				continue
			} else {
				unique_map[ustr] = true
			}
			s := v1 + v2
			two_sum_map[s] = append(two_sum_map[s], i*n+i+1+j)
		}
	}
	checked_number_map := make(map[int]bool)
	for k, v := range nums {
		_, checked := checked_number_map[v]
		if checked {
			continue
		} else {
			checked_number_map[v] = true
		}
		complete := 0 - v
		xl, ok := two_sum_map[complete]
		if ok {
			for _, x := range xl {
				if x/n <= k || x%n <= k {
					continue
				}
				l := []int{v, nums[x/n], nums[x%n]}
				sort.Ints(l)
				skip := false
				for m := 0; m < len(r); m++ {
					if r[m][0] < l[0] {
						continue
					} else if r[m][0] > l[0] {
						tmp := append([][]int{l}, r[m:]...)
						r = append(r[:m], tmp...)
						skip = true
						break
					}

					if r[m][1] < l[1] {
						continue
					} else if r[m][1] > l[1] {
						tmp := append([][]int{l}, r[m:]...)
						r = append(r[:m], tmp...)
						skip = true
						break
					}

					if r[m][2] < l[2] {
						continue
					} else if r[m][2] > l[2] {
						tmp := append([][]int{l}, r[m:]...)
						r = append(r[:m], tmp...)
						skip = true
						break
					}
					skip = true
					break
				}
				if !skip {
					r = append(r, l)
				}
			}
		}
	}
	return r
}
