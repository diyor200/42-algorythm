package main

type Solution struct{}

type Node struct {
	Val  int
	Next *Node
}

func (Solution) MiddleNode(head *Node) *Node {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
