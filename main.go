package main

// url: https://leetcode.com/problems/find-pivot-index/description/?envType=problem-list-v2&envId=prefix-sum

func pivotIndex(nums []int) int {
	rightSum, leftSum := 0, 0
	for i := range nums {
		rightSum += nums[i]
	}

	for i := range nums {
		if leftSum == rightSum-nums[i] {
			return i
		}
		leftSum += nums[i]
		rightSum -= nums[i]
	}

	return -1
}
