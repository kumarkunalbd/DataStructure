/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package l1120

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maximumAverageSubtree(root *TreeNode) float64 {
	_, _, ans := maxAvgTree(root)

	return ans
}

func maxAvgTree(root *TreeNode) (int, int, float64) {
	// base case
	if root == nil {
		return 0, 0, float64(0)
	}
	// main logic
	sL, nL, maxAvgL := maxAvgTree(root.Left)
	sR, nR, maxAvgR := maxAvgTree(root.Right)

	curSum := (sL + sR + root.Val)
	curN := nL + nR + 1
	maxAvg := float64(curSum) / float64(curN)

	if maxAvgL > maxAvg {
		maxAvg = maxAvgL
	}
	if maxAvgR > maxAvg {
		maxAvg = maxAvgR
	}
	return curSum, curN, maxAvg
}
