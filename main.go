package main

type Solution struct{}

func (Solution) LeftRightDifference(nums []int) []int {
	leftSum := 0
	rightSum := 0
	result := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			rightSum = sum(nums) - nums[i]
		} else {
			leftSum += nums[i-1]
			rightSum -= nums[i]
		}

		if leftSum > rightSum {
			result[i] = -1
		} else if leftSum < rightSum {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	return result
}

func sum(nums []int) int {
	result := 0
	for i := range nums {
		result += nums[i]
	}

	return result
}
