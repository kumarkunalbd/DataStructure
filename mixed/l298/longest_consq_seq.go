/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package l298

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestConsecutive(root *TreeNode) int {
	max := 0
	dfs(root, math.MinInt, 0, &max)
	return max
}

func dfs(root *TreeNode, target, curSeq int, maxP *int) {
	// base case
	if root == nil {
		return
	}
	// main logic
	if (*root).Val != target {
		curSeq = 0
	}
	curSeq++
	fmt.Printf("l1: node=%v target=%v curSeq=%v max=%v\n", (*root).Val, target, curSeq, *maxP)
	if curSeq > (*maxP) {
		*maxP = curSeq
	}
	fmt.Printf("l2: node=%v target=%v curSeq=%v max=%v\n", (*root).Val, target, curSeq, *maxP)

	dfs((*root).Left, (*root).Val+1, curSeq, maxP)
	dfs((*root).Right, (*root).Val+1, curSeq, maxP)

}
