package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	arr := []int{1, 2, 3, 4, 5}
	head := s.Reverse(createLinkedList(arr))

	assert.Equal(t, []int{5, 4, 3, 2, 1}, read(head))
	t.Log("pass")
}

func createLinkedList(arr []int) *Node {
	head := &Node{Val: arr[0]}
	current := head

	for _, v := range arr[1:] {
		current.Next = &Node{Val: v}
		current = current.Next
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
