/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) int {
	return dfs(root, math.MinInt, 0)
}

func dfs(root *TreeNode, target, curSeq int) int {
	// base case
	if root == nil {
		return 0
	}
	// main logic
	if (*root).Val != target {
		curSeq = 0
	}
	curSeq++
	lMax := dfs((*root).Left, (*root).Val+1, curSeq)
	rMax := dfs((*root).Right, (*root).Val+1, curSeq)

	if lMax > curSeq {
		curSeq = lMax
	}
	if rMax > curSeq {
		curSeq = rMax
	}

	return curSeq
}
