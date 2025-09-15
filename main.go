package main

type Solution struct{}

func (Solution) SortedSquares(nums []int) []int {
	n := len(nums)
	i, j := 0, n-1
	pos := n - 1
	var result = make([]int, n)

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
