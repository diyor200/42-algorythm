package main

type Solution struct{}

type Node struct {
	Val  int
	Next *Node
}

func (Solution) Reverse(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
