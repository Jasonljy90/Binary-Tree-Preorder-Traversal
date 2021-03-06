package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	helper(&res, root)
	return res
}

func helper(res *[]int, root *TreeNode) {
	if root.Left != nil {
		helper(res, root.Left)
	}
	if root.Right != nil {
		helper(res, root.Right)
	}
	*res = append(*res, root.Val)
}
