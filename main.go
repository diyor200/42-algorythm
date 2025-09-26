package main

type Solution struct{}

func (Solution) TwoSum(nums []int, target int) []int {
	set := make(map[int]struct{}, len(nums))
	res := make([]int, 2)

	for i := range nums {
		if _, ok := set[nums[i]]; ok {
			continue
		}

		set[nums[i]] = struct{}{}
	}

	return res
}
