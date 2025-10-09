package main

type Solution struct{}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (s Solution) LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		layer := []int{}
		temp := []*TreeNode{}
		for i := range queue {
			layer = append(layer, queue[i].Val)
			if queue[i].Left != nil {
				temp = append(temp, queue[i].Left)
			}
			if queue[i].Right != nil {
				temp = append(temp, queue[i].Right)
			}
		}
		res = append(res, layer)
		queue = temp
	}

	return res
}
