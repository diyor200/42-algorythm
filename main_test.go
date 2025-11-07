package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {

	testCases := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "first",
			nums1: []int{1, 2, 2, 2},
			nums2: []int{2, 2},
			want:  []int{2, 2},
		},
		{
			name:  "second",
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{4, 9},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, intersect(tt.nums1, tt.nums2))
		})
	}

	t.Log("pass")
}
