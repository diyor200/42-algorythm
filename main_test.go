package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{name: "first", nums: []int{1, 2, 3, 1}, k: 3, want: true},
		{name: "second", nums: []int{1, 0, 1, 1}, k: 1, want: true},
		{name: "third", nums: []int{1, 2, 3, 1, 2, 3}, k: 2, want: false},
		{name: "fourth", nums: []int{1, 2, 1}, k: 0, want: false},
		{name: "fifth", nums: []int{1}, k: 1, want: false},
		{name: "sixth", nums: []int{99, 99}, k: 2, want: true},
		{name: "seventh", nums: []int{99, 99}, k: 3, want: true},
		{name: "eights", nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, k: 15, want: false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, containsNearbyDuplicate(tt.nums, tt.k))
		})
	}

	t.Log("pass")
}
