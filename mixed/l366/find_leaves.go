package l366

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findLeaves(root *TreeNode) [][]int {
	ans := make([][]int, 0)

	for root.Left != nil || root.Right != nil {
		leaves := findAndDeleteLeaves(root, nil, false, make([]int, 0))
		ans = append(ans, leaves)
	}
	ans = append(ans, []int{root.Val})
	return ans

}

func findAndDeleteLeaves(root, parent *TreeNode, left bool, nodes []int) []int {
	// base case
	if root == nil {
		return nodes
	}
	if root.Left == nil && root.Right == nil {
		nodes = append(nodes, root.Val)
		if parent != nil && left {
			parent.Left = nil
		} else if parent != nil && !left {
			parent.Right = nil
		}
		return nodes
	}
	// cmd logic
	nodes = findAndDeleteLeaves(root.Left, root, true, nodes)
	nodes = findAndDeleteLeaves(root.Right, root, false, nodes)
	return nodes
}
