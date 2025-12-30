package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		nums []int
		want int
	}{
		{name: "first", nums: []int{1, 7, 3, 6, 5, 6}, want: 3},
		{name: "second", nums: []int{1, 2, 3}, want: -1},
		{name: "third", nums: []int{2, 1, -1}, want: 0},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, pivotIndex(tt.nums))
		})
	}

	t.Log("pass")
}
