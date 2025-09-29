package main

type Solution struct{}

func (Solution) ProductExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}

	return res
}
