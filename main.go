package main

type Solution struct{}

func (Solution) Rotate(nums []int, k int) []int {
	k = k % len(nums)
	l := len(nums)

	nums = reverse(nums, 0, l-1)
	nums = reverse(nums, k, l-1)
	nums = reverse(nums, 0, k-1)

	return nums
}

func reverse(nums []int, i, j int) []int {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	return nums
}
