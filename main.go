package main

// url: https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/description/

func minStartValue(nums []int) int {
	var startValue int
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if sum < startValue {
			startValue = sum
		}
	}

	if startValue < 0 {
		return -1*startValue + 1
	}

	return 1
}
