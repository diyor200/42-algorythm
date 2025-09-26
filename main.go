package main

type Solution struct{}

func (Solution) TwoSum(nums []int, target int) []int {
	compliments := make(map[int]int, len(nums))
	res := []int{}

	for i := range nums {
		if v, ok := compliments[nums[i]]; ok {
			res = append(res, v, i)
			break
		}

		compliments[target-nums[i]] = i
	}

	return res
}
