package main

import "testing"

func TestSolution(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	s := Solution{}
	t.Log(s.Rotate(nums, k))
}
