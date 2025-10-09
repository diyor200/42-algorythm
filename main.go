package main

type Solution struct{}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// steps:
// 1. invert right subtree of root
// 2. compare with left subtree
// func (s Solution) IsSymmetric(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	var queue = []*TreeNode{root.Left, root.Right}
// 	for len(queue) > 0 {
// 		left := queue[0]
// 		right := queue[1]
// 		queue = queue[2:]

// 		if left == nil && right == nil {
// 			continue
// 		}

// 		if left == nil || right == nil {
// 			return false
// 		}

// 		if left.Val != right.Val {
// 			return false
// 		}

// 		queue = append(queue, left.Left, right.Right)
// 		queue = append(queue, left.Right, right.Left)
// 	}

// 	return true
// }

// func (s Solution) IsSymmetric(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	queue := []*TreeNode{root.Left, root.Right}

// 	for len(queue) > 0 {
// 		left := queue[0]
// 		right := queue[1]
// 		queue = queue[2:]

// 		if left == nil && right == nil {
// 			continue
// 		}
// 		if left == nil || right == nil {
// 			return false
// 		}
// 		if left.Val != right.Val {
// 			return false
// 		}

// 		// Add mirrored children
// 		queue = append(queue, left.Left, right.Right)
// 		queue = append(queue, left.Right, right.Left)
// 	}

// 	return true
// }

func (s Solution) IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isEqual(root.Left, root.Right)
}

func isEqual(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isEqual(left.Left, right.Right) && isEqual(left.Right, right.Left)
}
