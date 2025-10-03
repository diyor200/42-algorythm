package main

type Solution struct{}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (Solution) Count_leaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var stack []*TreeNode
	var visited = make(map[*TreeNode]struct{})
	var res int

	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := visited[node]; !ok {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right == nil && node.Left == nil {
				res++
			}
			visited[node] = struct{}{}
		}
	}

	return res
}

// func (s Solution) Count_leaves(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return 1
// 	}

// 	return s.Count_leaves(root.Left) + s.Count_leaves(root.Right)
// }
