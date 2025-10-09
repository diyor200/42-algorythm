package main

type Solution struct{}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// diameter = radius + radius
func (s Solution) DiameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	findDepth(root, &maxDiameter)
	return maxDiameter
}

func findDepth(root *TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}

	left := findDepth(root.Left, maxDiameter)
	right := findDepth(root.Right, maxDiameter)

	// Update the maximum diameter if current path is longer
	if left+right > *maxDiameter {
		*maxDiameter = left + right
	}

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
