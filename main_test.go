package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "first", nums: []int{1, 2, 2, 4}, want: []int{2, 3}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.)
		})
	}

	t.Log("pass")
}
