package main

import "fmt"

type Solution struct{}

// url: https://leetcode.com/problems/missing-number/?envType=problem-list-v2&envId=bit-manipulation

func missingNumber(nums []int) int {
	n := len(nums)
	fmt.Println(n)
	s := (1 + n) * n / 2 // arithmetic progression
	for i := range nums {
		s -= nums[i]
		fmt.Println(s)
	}

	return s
}

func missingNumber(nums []int) int {
	xorNums := 0
	for i := range nums {
		xorNums ^= nums[i] ^ (i + 1)
	}

	return xorNums
}

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3}))
}
