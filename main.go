package main

type Solution struct{}

type Node struct {
	Val  int
	Next *Node
}

func (Solution) GetIntersectionNode(headA *Node, headB *Node) *Node {
	pA, pB := headA, headB

	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}

		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}
