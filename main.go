package main

// url: https://leetcode.com/problems/minimum-value-to-get-positive-step-by-step-sum/?envType=problem-list-v2&envId=prefix-sum

func minStartValue(nums []int) int {
	prefixInt := 0
	minv := 1
	for i := range nums {
		prefixInt += nums[i]
		if prefixInt < minv {
			minv = prefixInt
		}
	}
	if minv <= 0 {
		return 1 + (-1 * minv)
	}

	return minv
}
