package threeSum

import "sort"
import "strconv"

func ThreeSum(nums []int) [][]int {
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

func ThreeSum2(nums []int) [][]int {
	rmap := make(map[int]map[int]map[int]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				v1 := nums[i]
				v2 := nums[j]
				v3 := nums[k]
				if v1+v2+v3 != 0 {
					continue
				}
				if v1 > v2 {
					tmp := v1
					v1 = v2
					v2 = tmp
				}
				if v1 > v3 {
					tmp := v1
					v1 = v3
					v3 = tmp
				}
				if v2 > v3 {
					tmp := v2
					v2 = v3
					v3 = tmp
				}
				rmm, ok1 := rmap[v1]
				if !ok1 {
					rmm = make(map[int]map[int]bool)
					rmap[v1] = rmm
				}
				rmmm, ok2 := rmm[v2]
				if !ok2 {
					rmmm = make(map[int]bool)
					rmm[v2] = rmmm
				}
				_, ok3 := rmmm[v3]
				if !ok3 {
					rmmm[v3] = true
				}
			}
		}
	}
	r := [][]int{}

	key1 := []int{}
	for k := range rmap {
		key1 = append(key1, k)
	}
	sort.Ints(key1)
	for _, k1 := range key1 {
		key2 := []int{}
		for k := range rmap[k1] {
			key2 = append(key2, k)
		}
		sort.Ints(key2)
		for _, k2 := range key2 {
			key3 := []int{}
			for k := range rmap[k1][k2] {
				key3 = append(key3, k)
			}
			sort.Ints(key3)
			for _, k3 := range key3 {
				r = append(r, []int{k1, k2, k3})
			}
		}
	}
	return r
}
