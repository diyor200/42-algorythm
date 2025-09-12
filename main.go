package main

func (_ Solution) MoveZeroes(nums []int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			for j := i; j < l-1; j++ {
				nums[j] = nums[j+1]
			}
			nums[l-1] = 0
		}
	}

	return nums
}

type Solution struct{}

func (_ Solution) OptimizedMoveZeros(nums []int) []int {
	n := len(nums)
	pos := 0

	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[pos] = nums[i]
			pos++
		}
	}

	for pos < n {
		nums[pos] = 0
		pos++
	}

	return nums
}
