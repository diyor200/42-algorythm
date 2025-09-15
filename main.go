package main

type Solution struct{}

type Node struct {
	Val  int
	Next *Node
}

func (Solution) RemoveNthFromEnd(head *Node, n int) *Node {
	node := &Node{Next: head}
	fast, slow := node, node

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return node.Next
}
