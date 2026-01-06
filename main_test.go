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
		{name: "first", nums: []int{-3, 2, -3, 4, 2}, want: 5},
		{name: "second", nums: []int{1, 2}, want: 1},
		{name: "third", nums: []int{1, -2, -3}, want: 5},
		{name: "fourth", nums: []int{1, -1}, want: 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minStartValue(tt.nums))
		})
	}

	t.Log("pass")
}
