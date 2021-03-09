package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	ans := 0
	if root == nil {
		return 0
	}

	if root.Val >= low && root.Val <= high {
		ans = root.Val
	}
	ans += rangeSumBST(root.Left, low, high)
	ans += rangeSumBST(root.Right, low, high)

	return ans
}