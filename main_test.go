package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}
	arr := []int{-4, -1, 0, 3, 10}
	assert.Equal(t, []int{0, 1, 9, 16, 100}, s.SortedSquares(arr))
	t.Log("pass")
}
