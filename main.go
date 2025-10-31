package main

// link: https://leetcode.com/problems/range-sum-query-immutable/description/?envType=problem-list-v2&envId=array

type NumArray struct {
	Arr []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)

	prefSum := make([]int, n+1)
	prefSum[1] = nums[0]
	for i := 1; i < n; i++ {
		prefSum[i+1] = prefSum[i] + nums[i]
	}

	return NumArray{Arr: prefSum}
}

func (num *NumArray) SumRange(left int, right int) int {
	return num.Arr[right+1] - num.Arr[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
