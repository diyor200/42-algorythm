package main

type Solution struct{}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// recursive dfs solution
// T: O(n)
// S: O(h) -> best case: O(logn), worth case: O(n). But recursion uses call stack, that's why
// stack-based solution is better.
//
// func (s Solution) InvertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	root.Right, root.Left = root.Left, root.Right
// 	s.InvertTree(root.Right)
// 	s.InvertTree(root.Left)

// 	return root
// }

// stack-base dfs solution
// T: O(n)
// S: best case: O(h) -> O(logn), worth case: O(n)
//
// func (s Solution) InvertTree(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	var stack []*TreeNode
// 	stack = append(stack, root)

// 	for len(stack) > 0 {
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		if node == nil {
// 			continue
// 		}

// 		node.Left, node.Right = node.Right, node.Left
// 		stack = append(stack, node.Left, node.Right)
// 	}

// 	return root
// }

// bfs solution
func (s Solution) InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var queue []*TreeNode
	queue = append(queue, root)

	i := 0
	for len(queue) > i {
		node := queue[i]

		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		i++
	}

	return root
}
