package main

import "testing"

func TestMoveZero(t *testing.T) {
	s := Solution{}
	nums := []int{0, 1, 0, 3, 12}
	s.OptimizedMoveZeros(nums)
	t.Log(nums)
}
