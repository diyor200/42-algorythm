package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	n := len(nums)
	root := &TreeNode{Val: nums[n/2]}
	arrToBts(nums[:n/2], root.Left)
	arrToBts(nums[n/2:], root.Right)
	return root
}
func arrToBts(sl []int, node *TreeNode) {
	if len(sl) == 0 {
		return
	}

	if len(sl) == 1 {
		node.Left = &TreeNode{Val: sl[0]}
		return
	}

	node.Left = &TreeNode{Val: sl[0]}
	node.Right = &TreeNode{Val: sl[1]}
	arrToBts(sl[2:], node.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
