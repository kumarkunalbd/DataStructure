/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) int {
	_, longest := longestConseSeq(root)
	return longest
}

func longestConseSeq(root *TreeNode) (int, int) {
	// base case
	if root == nil {
		return 0, 0
	}

	// main logic
	curSeqL, longestL := longestConseSeq((*root).Left)
	if (*root).Left != nil && (*(*root).Left).Val-(*root).Val == 1 {
		curSeqL++
	} else {
		curSeqL = 1
	}
	curSeqR, longestR := longestConseSeq((*root).Right)
	if (*root).Right != nil && (*(*root).Right).Val-(*root).Val == 1 {
		curSeqR++
	} else {
		curSeqR = 1
	}
	maxSeq := int(math.Max(float64(curSeqL), float64(curSeqR)))
	longest := int(math.Max(float64(longestL), float64(longestR)))
	if maxSeq > longest {
		longest = maxSeq
	}

	return maxSeq, longest

}