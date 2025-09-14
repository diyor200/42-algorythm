package main

import "testing"

func TestSolution(t *testing.T) {
	// input := "((()))"
	s := Solution{}
	// t.Log(s.Is_valid(input))

	input1 := "([)"
	t.Log(s.isValid(input1))
}
