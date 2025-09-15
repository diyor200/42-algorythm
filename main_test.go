package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	arr := []int{1, 2, 2, 1}
	assert.Equal(t, true, s.IsPalindrome(create(arr)))
	t.Log("pass")
}

func create(arr []int) *Node {
	head := &Node{Val: arr[0]}
	curr := head

	for _, v := range arr[1:] {
		curr.Next = &Node{Val: v}
		curr = curr.Next
	}

	return head
}

func read(head *Node) []int {
	var res []int

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}
