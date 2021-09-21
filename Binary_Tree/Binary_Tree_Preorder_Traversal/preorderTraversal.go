package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	traversal := make([]int, 0, 1000)
	if root == nil {
		return traversal
	}
	node := root

	breadcrumbs := make([]*TreeNode, 0, 1000)
	var prev_node *TreeNode
	var reverse bool
	reverse = false

	for i := 0; ; i++ {
		if !reverse {
			breadcrumbs = append(breadcrumbs, node)
			traversal = append(traversal, node.Val)
		}
		// fmt.Printf("Node address: %p\n", node)
		// fmt.Println("    i:" , i, "node:", node, "address: ", &node, "traversal:", traversal, "breadcrumbs:", breadcrumbs)
		// fmt.Println("    reverse:" , reverse)
		if node.Left != nil && !reverse {
			node = node.Left
			reverse = false
			continue
		}
		if node.Right != nil && prev_node != node.Right {
			node = node.Right
			prev_node = nil
			reverse = false
			continue
		}
		if len(breadcrumbs) <= 1 {
			break
		}
		breadcrumbs = breadcrumbs[:len(breadcrumbs)-1]
		prev_node = node
		node = breadcrumbs[len(breadcrumbs)-1]
		reverse = true
	}
	return traversal
}
