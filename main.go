package main

import "slices"

func findErrorNums(nums []int) []int {
	res := make([]int, 2)
	slices.Sort(nums)
	numMap := make(map[int]bool, len(nums))

	for i := range nums {
		if res[1] == 0 && nums[i] != (i+1) {
			res[1] = i + 1
		}
		if _, ok := numMap[nums[i]]; ok {
			res[0] = nums[i]
		}

		numMap[nums[i]] = true
	}

	return res
}
