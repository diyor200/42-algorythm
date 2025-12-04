package main

import "fmt"

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	//list = &ListNode{Val: 1}

	head := rotateRight(list, 99)
	read(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	l := length(head)
	k = k % l

	if k == 0 {
		return head
	}

	node := head
	for i := 0; i < l-k-1; i++ {
		node = node.Next
	}

	newHead := node.Next
	tail := newHead
	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = head
	node.Next = nil
	return newHead
}

func length(node *ListNode) int {
	res := 0

	for node != nil {
		res++
		node = node.Next
	}

	return res
}

func read(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
