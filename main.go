package main

// url: https://leetcode.com/problems/shortest-distance-to-a-character/?envType=problem-list-v2&envId=two-pointers

func shortestToChar(s string, c byte) []int {
	n := len(s)
	res := make([]int, n)

	prev := -100_000
	for i := 0; i < n; i++ {
		if s[i] == c {
			prev = i
		}

		res[i] = i - prev
	}

	prev = 100_000
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			prev = i
		}

		if prev-i < res[i] {
			res[i] = prev - i
		}
	}

	return res
}
