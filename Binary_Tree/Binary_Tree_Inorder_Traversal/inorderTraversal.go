package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	// a := 4
	fmt.Println(root)
	traversal := make([]int, 0, 1000)
	breadcrumbs = make([]*TreeNode, 0, 1000)

	for root != nil {
		breadcrumbs = append(breadcrumbs, root)
		if root.Left == nil {
			traversal = append(append, root.Val)
		}
	}

	return traversal
}
