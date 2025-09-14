package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	l := []int{1, 2, 3, 4, 5}
	head := createLinkedList(l)

	middle := s.MiddleNode(head)
	t.Log(middle)
	read(middle)
}

func createLinkedList(list []int) *Node {
	head := &Node{Val: list[0]}
	current := head

	for _, v := range list[1:] {
		current.Next = &Node{Val: v}
		current = current.Next
	}

	return head

}

func read(head *Node) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
