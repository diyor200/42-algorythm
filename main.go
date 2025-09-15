package main

type Solution struct{}

func (Solution) SortedSquares(nums []int) []int {
	var result = make([]int, len(nums))
	i, j := 0, len(nums)-1
	pos := j

	for i <= j {
		leftSq := nums[i] * nums[i]
		rightSq := nums[j] * nums[j]

		if leftSq > rightSq {
			result[pos] = leftSq
			i++
		} else {
			result[pos] = rightSq
			j--
		}
		pos--
	}

	return result
}
