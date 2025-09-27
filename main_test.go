package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		name string
		arr  []string
		want int
	}{
		{name: "first", arr: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, want: 3},
		{name: "second", arr: []string{"a"}, want: 1},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, s.GroupAnagrams(tt.arr))
		})
	}

	t.Log("pass")
}

func TestAngramGroups(t *testing.T) {
	testCases := []struct {
		name string
		arr  []string
		want [][]string
	}{
		{name: "first", arr: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{name: "second", arr: []string{""}, want: [][]string{{""}}},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, groupAnagrams(tt.arr))
		})
	}

	t.Log("pass")
}
