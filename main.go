package main

type Solution struct{}

// my first solution
func (Solution) MoveZeroes(nums []int) []int {
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

func (Solution) MoveZeroes42(nums []int) []int {
	zeroCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
			continue
		}

		nums[i], nums[i-zeroCount] = nums[i-zeroCount], nums[i]
	}

	return nums
}

func (Solution) OptimizedMoveZeros(nums []int) []int {
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
