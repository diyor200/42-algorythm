package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	arr1 := []int{1, 2, 3, 4, 5}
	n1 := 2
	assert.Equal(t, []int{1, 2, 3, 5}, read(s.RemoveNthFromEnd(create(arr1), n1)))

	arr2 := []int{1}
	n2 := 1
	assert.Equal(t, nil, read(s.RemoveNthFromEnd(create(arr2), n2)))

	t.Log("pass")
}

func create(arr []int) *Node {
	head := &Node{Val: arr[0]}
	curr := head

	for _, val := range arr[1:] {
		curr.Next = &Node{Val: val}
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
