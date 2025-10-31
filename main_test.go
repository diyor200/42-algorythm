package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name        string
		nums        []int
		left, right int
		want        int
	}{
		{name: "first", nums: []int{-2, 0, 3, -5, 2, -1}, want: 1, left: 0, right: 2},
		{name: "second", nums: []int{-2, 0, 3, -5, 2, -1}, want: -1, left: 2, right: 5},
		{name: "first", nums: []int{-2, 0, 3, -5, 2, -1}, want: -3, left: 0, right: 5},
	}

	for _, tt := range testCases {
		s := Constructor(tt.nums)
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.SumRange(tt.left, tt.right))
		})
	}

	t.Log("pass")
}
