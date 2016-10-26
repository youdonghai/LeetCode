package findMedianSortedArrays

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	all_length := m + n
	sorted_nums := make([]int, m+n, m+n)
	i1, i2 := 0, 0
	for i := 0; i < all_length; i++ {
		if i1 == m {
			for ; i2 < n; i2++ {
				sorted_nums[i] = nums2[i2]
				i++
			}
			break
		}
		if i2 == n {
			for ; i1 < m; i1++ {
				sorted_nums[i] = nums1[i1]
				i++
			}
			break
		}

		if nums1[i1] < nums2[i2] {
			sorted_nums[i] = nums1[i1]
			i1++
		} else if nums2[i2] < nums1[i1] {
			sorted_nums[i] = nums2[i2]
			i2++
		} else {
			sorted_nums[i] = nums1[i1]
			sorted_nums[i+1] = nums2[i2]
			i1++
			i2++
			i++
		}
	}
	all_length--
	return float64(sorted_nums[all_length/2]+sorted_nums[all_length/2+all_length%2]) / 2.0
}
