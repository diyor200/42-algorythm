package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		name  string
		trust [][]int
		n     int
		want  int
	}{
		{name: "first", n: 2, trust: [][]int{{1, 2}}, want: 2},
		{name: "second", n: 1, trust: [][]int{}, want: 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.FindJudge(tt.n, tt.trust))
		})
	}

	t.Log("pass")
}
