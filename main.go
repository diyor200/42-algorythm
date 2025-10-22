package main

type Solution struct{}

// url: https://leetcode.com/problems/counting-bits/?envType=problem-list-v2&envId=bit-manipulation

// T: O(nlogn)
func countBits(n int) []int {
	res := make([]int, n+1)

	for i := 1; i <= n; i++ {
		res[i] = countOnes(i)
	}

	return res
}

func countOnes(n int) int {
	count := 0
	for n != 0 {
		count++
		n &= (n - 1)
	}

	return count
}

// optimized: T O(n)

func optimized(n int) []int {
	var res = make([]int, n+1)

	for i := 1; i <= n; i++ {
		res[i] = res[i>>1] + (i & 1)
	}

	return res
}
