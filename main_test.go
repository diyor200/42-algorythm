package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSolution(t *testing.T) {
	s := Solution{}

	// case 1
	intersection := &Node{Val: 2, Next: &Node{Val: 4}}
	arr1 := []int{1, 9, 1}
	arr2 := []int{3}
	headA := create(arr1)
	headB := create(arr2)
	headA.Next = intersection
	headB.Next = intersection
	assert.Equal(t, 2, s.GetIntersectionNode(headA, headB).Val)

	//case 2
	intersection = &Node{Val: 8, Next: &Node{Val: 4, Next: &Node{Val: 5}}}
	arr1 = []int{4, 1}
	arr2 = []int{5, 6, 1}
	headA = create(arr1)
	headB = create(arr2)
	headA.Next = intersection
	headB.Next = intersection
	assert.Equal(t, 8, s.GetIntersectionNode(headA, headB).Val)

	arr1 = []int{2, 6, 4}
	arr2 = []int{1, 5}
	t.Log(s.GetIntersectionNode(create(arr1), create(arr2)))
	assert.Equal(t, nil, s.GetIntersectionNode(create(arr1), create(arr2)))

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
