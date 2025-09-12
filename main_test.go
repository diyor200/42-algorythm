package main

import "testing"

func TestSolution(t *testing.T) {
	s := []string{"h", "e", "l", "l", "o"}

	sol := Solution{}
	t.Log(sol.ReverseString(s))
}
