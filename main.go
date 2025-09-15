package main

type Solution struct{}

type Node struct {
	Val  int
	Next *Node
}

func findMiddle(head *Node) *Node {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverse(head *Node) *Node {
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

func isSame(head1, head2 *Node) bool {
	for head1 != nil && head2 != nil {
		if head1.Val != head2.Val {
			return false
		}

		head1 = head1.Next
		head2 = head2.Next
	}

	return true
}

func (Solution) IsPalindrome(head *Node) bool {
	middle := findMiddle(head)
	middle = reverse(middle)

	return isSame(head, middle)
}
