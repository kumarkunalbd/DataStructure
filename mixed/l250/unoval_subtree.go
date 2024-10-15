/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package l250

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	ans, _ := univalSubtrees(root)
	return ans
}

func univalSubtrees(root *TreeNode) (int, bool) {
	// base case
	if root == nil {
		return 0, true
	}

	if root.Left == nil && root.Right == nil {
		return 1, true
	}

	// cmd logic
	lVal, isLU := univalSubtrees((*root).Left)
	rVal, isRU := univalSubtrees((*root).Right)
	if !isLU || !isRU {
		return lVal + rVal, false
	}

	if root.Left != nil && root.Right == nil {
		if root.Val == root.Left.Val {
			return lVal + 1, true
		}
		return lVal, false
	}
	if root.Right != nil && root.Left == nil {
		if root.Right.Val == root.Val {
			return rVal + 1, true
		}
		return rVal, false
	}

	if root.Val == root.Left.Val && root.Val == root.Right.Val {
		return lVal + rVal + 1, true
	}
	return lVal + rVal, false
}
