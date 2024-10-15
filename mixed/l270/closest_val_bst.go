package l270

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestValue(root *TreeNode, target float64) int {
	ans, _ := closestValueWithDiff(root, target)

	return ans
}

func closestValueWithDiff(root *TreeNode, target float64) (int, float64) {
	// base case
	if root == nil {
		return math.MaxInt, math.MaxFloat64
	}
	// cmd logic
	curDiff := math.Abs(float64(root.Val) - target)
	closestVal := root.Val
	if float64(root.Val) > target {
		lClose, lDiff := closestValueWithDiff(root.Left, target)
		if lDiff <= curDiff {
			curDiff = lDiff
			closestVal = lClose
		}
	} else if float64(root.Val) < target {
		rClose, rDiff := closestValueWithDiff(root.Right, target)
		if rDiff < curDiff {
			curDiff = rDiff
			closestVal = rClose
		}
	}
	return closestVal, curDiff
}
