package main

func intersect(nums1 []int, nums2 []int) []int {
	var (
		res []int
		set = make(map[int]int)
	)

	for i := range nums1 {
		set[nums1[i]]++
	}

	for i := range nums2 {
		if _, ok := set[nums2[i]]; ok {
			if set[nums2[i]] > 0 {
				res = append(res, nums2[i])
				set[nums2[i]]--
			}
		}
	}

	return res
}
